package paymenthandler

import (
	"fmt"
    "math/rand"
    "time"
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/coreapi"
    "github.com/midtrans/midtrans-go/snap"
    "os"
    // "github.com/midtrans/midtrans-go/iris"
)

func GenerateOrderID() string {
    rand.Seed(time.Now().UnixNano()) // Set seed dari waktu sekarang agar randomizer selalu berubah

    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Set karakter yang mungkin digunakan pada order ID
    const idLength = 10                                       // Panjang order ID yang diinginkan

    var result string
    for i := 0; i < idLength; i++ {
        randomChar := string(charset[rand.Intn(len(charset))]) // Ambil karakter random dari charset
        result += randomChar                                   // Tambahkan karakter tersebut ke dalam order ID
    }

    return result
}


func GenerateRandomQRCode() string {
    // Seed generator acak dengan nilai waktu saat ini
    rand.Seed(time.Now().UnixNano())

    // Karakter yang diizinkan untuk QR code
    allowedChars := "ZETA-ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // Panjang string QR code yang diinginkan
    qrCodeLength := 20

    // Buat sebuah string acak dengan karakter yang diizinkan
    qrCode := make([]byte, qrCodeLength)
    for i := range qrCode {
        qrCode[i] = allowedChars[rand.Intn(len(allowedChars))]
    }

    // Kembalikan string QR code acak
    return string(qrCode)
}

func GenerateTicketID() string {
    // Format tiket ID: TKT-xxxx-xxxx-xxxx-xxxx
    // xxxx adalah bilangan acak empat digit atau karakter acak empat digit
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 4)
    rand.Read(b)
    randomStr1 := fmt.Sprintf("%x", b)

    b = make([]byte, 4)
    rand.Read(b)
    randomStr2 := fmt.Sprintf("%x", b)

    b = make([]byte, 4)
    rand.Read(b)
    randomStr3 := fmt.Sprintf("%x", b)

    b = make([]byte, 4)
    for i := range b {
        b[i] = byte(rand.Intn(10) + 48)
    }
    randomStr4 := string(b)

    return fmt.Sprintf("TKT-%s-%s-%s-%s", randomStr1, randomStr2, randomStr3, randomStr4)
}

func CheckTransactionStatus(transactionID string) string {
    // Konfigurasi kredensial Midtrans
    // midclient := midtrans.NewClient()
    // midclient.ServerKey = "SB-Mid-server-LqlhIR_ldHS0kfZkcfo71dLO"
    // midclient.ClientKey = "SB-Mid-client-vmpBhhQpamztBLyj"
	var c coreapi.Client
    c.New(os.Getenv("SNB_M_SERVERKEY"), midtrans.Sandbox)

	res, err := c.CheckTransaction(transactionID)
	if err != nil {
		// do something on error handle
        return "transaction not Found!"
	}
	return res.TransactionID

    // Jika tidak terjadi kesalahan, kembalikan objek TransactionStatusResponse
    
}

func SnapTranc(fname, email, phone, orderid string, amount int64) (string, error) {
    midtrans.ServerKey = os.Getenv("SNB_M_SERVERKEY")
	midtrans.ClientKey = os.Getenv("SNB_M_CLIENTKEY")
	midtrans.Environment = midtrans.Sandbox
    // Set transaction details
	// amount := "100000" // in rupiah
	customerDetails := &midtrans.CustomerDetails{
		FName: fname,
		Email: email,
		Phone: phone,
	}

	// Set Snap payment method specific parameters

	// Set Snap transaction details
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderid,
			GrossAmt: amount,
		},
		CustomerDetail: customerDetails,
		// EnabledPayments: snap.AllSnapPaymentType,
        EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeBCAVA,
			snap.PaymentTypeBNIVA,
			snap.PaymentTypeBRIVA,
			snap.PaymentTypeOtherVA,
			snap.PaymentTypeIndomaret,
			snap.PaymentTypeAlfamart,
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
			snap.PaymentTypeOtherVA,
		},
        Callbacks: &snap.Callbacks{
			Finish: "https://www.youtube.com/watch?v=5Bkub_BYZvU",
		},
	}

	// Create Snap transaction token
	tokenResp, err := snap.CreateTransaction(snapReq)
	if err != nil {
		return "", fmt.Errorf("error creating Snap transaction token: %v", err)
	}

	return tokenResp.RedirectURL, nil
}

// func main() {
    // BtPayment("order-101", 44000, "bank_transfer", "bca")
	// CheckTransactionStatus("ZKCULFPN5I")
	// orderids := generateOrderID()
	// fmt.Println(orderids)
    // fmt.Println(SnapTranc("megawati","mega@gmail.com","0857710149741",orderids))
// }