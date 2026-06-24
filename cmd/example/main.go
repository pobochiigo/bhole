package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	agencyclient "github.com/pobochiigo/bhole/client/agency"
	launchclient "github.com/pobochiigo/bhole/client/launch"
	"github.com/pobochiigo/bhole/client/transport"
	"github.com/pobochiigo/bhole/internal/agency"
	"github.com/pobochiigo/bhole/internal/launch"
)

func main() {
	baseURL := "https://lldev.thespacedevs.com"
	fmt.Printf("Initializing LL2 ConnectRPC-to-REST client for %s...\n", baseURL)

	// Create custom HTTP client transport
	rawClient := &http.Client{Timeout: 30 * time.Second}
	restClient := transport.NewRESTClient(baseURL, rawClient)

	// Initialize the business services using our ConnectRPC client constructors
	launchesService := launchclient.NewLaunchClient(restClient, baseURL)
	agenciesService := agencyclient.NewAgencyClient(restClient, baseURL)

	ctx := context.Background()

	// 1. List launches
	fmt.Println("\n--- Listing Launches ---")
	launchesResp, err := launchesService.ListLaunches(ctx, &launch.ListLaunchesRequest{
		Limit: 2,
	})
	if err != nil {
		log.Fatalf("Error listing launches: %v", err)
	}

	fmt.Printf("Total launches count: %d\n", launchesResp.Count)
	for i, l := range launchesResp.Results {
		fmt.Printf("Launch %d:\n", i+1)
		fmt.Printf("  ID:   %s\n", l.Id)
		fmt.Printf("  Name: %s\n", l.Name)
		fmt.Printf("  Net:  %s\n", l.Net)
		if l.Status != nil {
			fmt.Printf("  Status: %s (%s)\n", l.Status.Name, l.Status.Abbrev)
		}
	}

	// 2. Get a single launch by ID if results are returned
	if len(launchesResp.Results) > 0 {
		targetID := launchesResp.Results[0].Id
		fmt.Printf("\n--- Getting Launch by ID (%s) ---\n", targetID)
		singleLaunch, err := launchesService.GetLaunch(ctx, &launch.GetLaunchRequest{
			ID: targetID,
		})
		if err != nil {
			log.Fatalf("Error getting launch: %v", err)
		}
		fmt.Printf("Launch Detail:\n")
		fmt.Printf("  Name:        %s\n", singleLaunch.Name)
		fmt.Printf("  Net:         %s\n", singleLaunch.Net)
		fmt.Printf("  Window Start: %s\n", singleLaunch.WindowStart)
		if singleLaunch.Status != nil {
			fmt.Printf("  Description: %s\n", singleLaunch.Status.Description)
		}
	}

	// 3. List agencies
	fmt.Println("\n--- Listing Agencies ---")
	agenciesResp, err := agenciesService.ListAgencies(ctx, &agency.ListAgenciesRequest{
		Limit: 2,
	})
	if err != nil {
		log.Fatalf("Error listing agencies: %v", err)
	}

	fmt.Printf("Total agencies count: %d\n", agenciesResp.Count)
	for i, a := range agenciesResp.Results {
		fmt.Printf("Agency %d:\n", i+1)
		fmt.Printf("  ID:     %d\n", a.Id)
		fmt.Printf("  Name:   %s (%s)\n", a.Name, a.Abbrev)
		if a.TypeVal != nil {
			fmt.Printf("  Type:   %s\n", a.TypeVal.Name)
		}
	}

	// 4. Get a single agency by ID if results are returned
	if len(agenciesResp.Results) > 0 {
		targetID := agenciesResp.Results[0].Id
		fmt.Printf("\n--- Getting Agency by ID (%d) ---\n", targetID)
		singleAgency, err := agenciesService.GetAgency(ctx, &agency.GetAgencyRequest{
			ID: targetID,
		})
		if err != nil {
			log.Fatalf("Error getting agency: %v", err)
		}
		fmt.Printf("Agency Detail:\n")
		fmt.Printf("  Name:          %s\n", singleAgency.Name)
		if singleAgency.Administrator != nil {
			fmt.Printf("  Administrator: %s\n", *singleAgency.Administrator)
		}
		if singleAgency.FoundingYear != nil {
			fmt.Printf("  Founding Year: %d\n", *singleAgency.FoundingYear)
		}
	}
}
