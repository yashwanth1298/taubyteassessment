package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/base64"
  "fmt"
  "io"
  "strings"
   "os"
   "bytes"

    shell "github.com/ipfs/go-ipfs-api"
)

type pairs struct {
	Data string `json:"data"`
	Key string `json:"key"`
	Cid string `json:"cid"`
}


func main() {
  // Echo instance
  e := echo.New()

  e.Use(middleware.CORS())

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // New Routes
  e.POST("/add", encryptAndStore)
  e.GET("/cid/:cid/:key", getDecryptedCidContents)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func getDecryptedCidContents(c echo.Context) error {
	cid := c.Param("cid")
	key := c.Param("key")
	
	cidcontent := getCidContents(cid)

        new_course := new(pairs)
        if err := c.Bind(new_course); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        fmt.Println("new pair %s", new_course.Data)
        fmt.Println("new pair %s", new_course.Key)

        encr_course := new(pairs)
        bytedata := []byte(key)
        encr_course.Data = decrypt(bytedata, cidcontent)
        encr_course.Key = key

        return c.JSON(http.StatusOK, encr_course.Data)
}

func encryptAndStore(c echo.Context) error {
	new_course := new(pairs)
	if err := c.Bind(new_course); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	fmt.Println("new pair %s", new_course.Data)
	fmt.Println("new pair %s", new_course.Key)

	encr_course := new(pairs)
    	bytedata := []byte(new_course.Key)
	encr_course.Data = encrypt(bytedata, new_course.Data)
 	encr_course.Key = new_course.Key

	cid := storeString(encr_course.Data)
        //fmt.Println("cid in encryptandstore function is %s", cid)
        encr_course.Cid = cid

        fmt.Println("encr_course in encryptandstore function is %s", encr_course)

	return c.JSON(http.StatusCreated, encr_course)
}


// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

//Store string
func storeString(mystr string) string {
        // Where your local node is running on localhost:5001
        sh := shell.NewShell("localhost:5001")

        cid, err := sh.Add(strings.NewReader(mystr))
        if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
        }
        fmt.Printf("Added \"%s\" in cid \"%s\" \n", mystr, cid)
        return cid
}

//To get contents of the cid
func getCidContents(cid string) string {
        // Where your local node is running on localhost:5001
        sh := shell.NewShell("localhost:5001")

        data, err := sh.Cat(cid)
        if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
        }

        buf := new(bytes.Buffer)
        buf.ReadFrom(data)
        newStr := buf.String()
        fmt.Printf("Contents of cid \"%s\":\n%s \n", cid, newStr)

	return newStr
}

