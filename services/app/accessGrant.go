package app

import (
	"context"
	"log"
	"os"

	"storj.io/uplink"
)

func AccessGrant(ctx context.Context, satelliteAddress, apiKey, appPassPhrase string) error { 
	access, err := uplink.RequestAccessWithPassphrase(ctx, satelliteAddress, apiKey, appPassPhrase)
	if err != nil {
		return err 
	}

	// Create an access grant for reading bucket "app"
	permission := uplink.ReadOnlyPermission()
	shared := uplink.SharePrefix{Bucket: "app"}
	restrictedAccess, err := access.Share(permission, shared)

	if err != nil {
		return err 
	}

	// serialize the restricted access grant
	serializedAccess, err := restrictedAccess.Serialize()
	if err != nil {
		return err
	}
	// fmt.Println(serializedAccess)

	os.Setenv("APPACCESS", serializedAccess)
	log.Println("Successfully generated app acess grant for the app bucket")
	return nil
}

