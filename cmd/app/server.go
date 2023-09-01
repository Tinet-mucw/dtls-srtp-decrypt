package app

import (
	// "github.com/Tinet-mucw/dtls-srtp-decrypt/pkg/dtls"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "[objectType]",
		Short: "[objectType]",
		Long: `[objectType].
objectType: wav, mp3`,
		Args: cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validationEncoder(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return masterSecretCal(args)
		},
	}

	return cmd
}

func masterSecretCal(args []string) error {
	klog.Infof("execute cal master secret")
	// dtls.CalPreMasterSecret()
	return nil
}

func validationEncoder(args []string) error {
	return nil
}
