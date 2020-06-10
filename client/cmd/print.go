package cmd

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	proto "github.com/thomas-maurice/thermal-printer/go"
)

var (
	pixelSize     int64
	centerQR      bool
	centerBarcode bool
	blankLines    int32
	font          int64
	addNewLine    bool
	barcodeHeight int64
	barcodeWidth  int64
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Send print jobs",
	Long:  ``,
}

var printQRCmd = &cobra.Command{
	Use:   "qr",
	Short: "Prints a QR code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatal("You should only pass the data to encode in the QR code")
		}
		c, err := getClient()
		if err != nil {
			logrus.WithError(err).Fatal("Could not get a client")
		}

		data, err := c.QR(
			context.Background(),
			&proto.QRCode{
				Code:      args[0],
				PixelSize: pixelSize,
				Center:    centerQR,
			},
		)

		if err != nil {
			logrus.WithError(err).Fatal("Call failed")
		}
		output(data)
	},
}

var printBarcodeCmd = &cobra.Command{
	Use:   "barcode",
	Short: "Prints a barcode",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatal("You should only pass the data to encode in the barcode")
		}
		c, err := getClient()
		if err != nil {
			logrus.WithError(err).Fatal("Could not get a client")
		}

		data, err := c.Bar(
			context.Background(),
			&proto.Barcode{
				Code:   args[0],
				Center: centerBarcode,
				Blanks: blankLines,
			},
		)

		if err != nil {
			logrus.WithError(err).Fatal("Call failed")
		}
		output(data)
	},
}

var printBlankCmd = &cobra.Command{
	Use:   "blank",
	Short: "Print blank lines",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			logrus.WithError(err).Fatal("Could not get a client")
		}

		data, err := c.Blank(
			context.Background(),
			&proto.BlankLines{
				Number: blankLines,
				Font:   font,
			},
		)

		if err != nil {
			logrus.WithError(err).Fatal("Call failed")
		}
		output(data)
	},
}

var printLineCmd = &cobra.Command{
	Use:   "line",
	Short: "Prints a line",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatal("You should only pass the data to print")
		}
		c, err := getClient()
		if err != nil {
			logrus.WithError(err).Fatal("Could not get a client")
		}

		data, err := c.Print(
			context.Background(),
			&proto.Line{
				Line: args[0],
				Font: font,
			},
		)
		if err != nil {
			logrus.WithError(err).Fatal("Call failed")
		}
		output(data)

		if addNewLine {
			data, err := c.Blank(
				context.Background(),
				&proto.BlankLines{
					Number: 1,
					Font:   font,
				},
			)

			if err != nil {
				logrus.WithError(err).Fatal("Call failed")
			}
			output(data)
		}
	},
}

func initPrintCmd() {
	printQRCmd.PersistentFlags().BoolVarP(&centerQR, "center", "", true, "Wether or not center the QR code")
	printQRCmd.PersistentFlags().Int64VarP(&pixelSize, "pixel-size", "p", 5, "Pixel size for the code")

	printBlankCmd.PersistentFlags().Int32VarP(&blankLines, "blank", "b", 1, "How many lines to output")
	printBlankCmd.PersistentFlags().Int64VarP(&font, "font", "f", 1, "Font to use")

	printLineCmd.PersistentFlags().Int64VarP(&font, "font", "f", 1, "Font to use")
	printLineCmd.PersistentFlags().BoolVarP(&addNewLine, "new-line", "n", true, "Wether or not add a new line")

	printBarcodeCmd.PersistentFlags().BoolVarP(&centerBarcode, "center", "", true, "Wether or not center the barcode")
	printBarcodeCmd.PersistentFlags().Int32VarP(&blankLines, "blank", "b", 1, "How many lines to output")

	printCmd.AddCommand(printLineCmd)
	printCmd.AddCommand(printQRCmd)
	printCmd.AddCommand(printBlankCmd)
	printCmd.AddCommand(printBarcodeCmd)
}
