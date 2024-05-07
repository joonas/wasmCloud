package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

const (
	KindStartComponent = "startComponent"
)

type PolicyRequest struct {
	RequestID string              `json:"requestId"`
	Kind      string              `json:"kind"`
	Version   string              `json:"version"`
	Request   PolicyRequestDetail `json:"request"`
}

type PolicyRequestDetail struct {
	ComponentID string `json:"componentId"`
	ProviderID  string `json:"providerId"`
	ImageRef    string `json:"imageRef"`
}

func (pr PolicyRequest) IsComponent() bool {
	return pr.Kind == KindStartComponent
}

type PolicyResponse struct {
	RequestId string `json:"requestId"`
	Permitted bool   `json:"permitted"`
	Message   string `json:"message,omitempty"`
}

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		natsUrl     = flags.String("nats-url", nats.DefaultURL, "NATS Server URL to connect to")
		policyTopic = flags.String("policy-topic", "wasmcloud.policy", "NATS Topic to listen for policy requests")
	)
	if err := flags.Parse(args[:1]); err != nil {
		return err
	}

	fmt.Printf("Attempting to connect to NATS on %s\n", *natsUrl)

	nc, err := nats.Connect(*natsUrl)
	if err != nil {
		fmt.Errorf("unable to connect to nats: %w\n", err)
	}
	defer nc.Drain()

	fmt.Printf("Attempting to listen on %s\n", *policyTopic)

	_, err = nc.Subscribe(*policyTopic, outputPolicyRequests)

	if err != nil {
		fmt.Errorf("unable to subscribe to policy topic (%q): %v\n", *policyTopic, err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	return nil
}

func outputPolicyRequests(m *nats.Msg) {
	var req PolicyRequest

	err := json.Unmarshal(m.Data, &req)
	if err != nil {
		fmt.Printf("Unable to marshall message: %v", err)
	}

	respId := uuid.New().String()
	if req.IsComponent() {
		fmt.Printf("Got component: %s\n", req.Request.ComponentID)
		b, artifactDigest, err := BundleFromOCIImage(req.Request.ImageRef)
		if err != nil {
			fmt.Printf("Unable to parse bundle: %w\n", err)

			resp, err := json.Marshal(PolicyResponse{
				RequestId: respId,
				Permitted: false,
				Message:   "Could not find an associated sigstore attestation.",
			})
			if err != nil {
				fmt.Printf("Failed to marshal response: %w", err)
				return
			}

			err = m.Respond(resp)
			if err != nil {
				fmt.Printf("Failed to respond to policy request: %w", err)
				return
			}
			return
		}
		fmt.Printf("	Artifact Digest: %s\n", *artifactDigest)
		fmt.Printf("	Bundle Mediatype: %s\n", b.GetMediaType())
		resp, err := json.Marshal(PolicyResponse{
			RequestId: respId,
			Permitted: true,
		})
		if err != nil {
			fmt.Printf("Failed to marshal response: %w", err)
			return
		}

		err = m.Respond(resp)
		if err != nil {
			fmt.Printf("Failed to respond to policy request: %w", err)
			return
		}

	} else {
		resp, err := json.Marshal(PolicyResponse{
			RequestId: respId,
			Permitted: true,
		})
		if err != nil {
			fmt.Printf("Failed to marshal response: %w", err)
		}

		err = m.Respond(resp)
		if err != nil {
			fmt.Printf("Failed to respond to policy request: %w", err)
		}
		return
	}

	fmt.Printf("Received a message: %s\n", string(m.Data))
}
