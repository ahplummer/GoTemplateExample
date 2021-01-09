package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"log"
	"math/rand"
	"time"
	"github.com/mitchellh/go-homedir"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func LambdaHandler(ctx context.Context) (Response, error) {
	html := UIDisplay()

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            html,
		Headers: map[string]string{
			"Content-Type":           "text/html",
		},
	}

	return resp, nil
}

func initConfig() (err error) {
	log.Printf("Starting GoExample, getting config")
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		return err
	}
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("~")
	viper.SetConfigName("maf.toml")
	pflag.Int("PORT", 8888, "This is the port to bind to. 8888 is default.")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Uh oh. no config file")
		err = nil
	}
	fmt.Printf("Will use %d as port number.\n", viper.GetInt("port"))
	if viper.GetInt("port") == 0 {
		err = errors.New("No valid port number")
	}
	return err
}

func NewRootHandler(w http.ResponseWriter, r *http.Request) {

	html := UIDisplay()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, html)

}

func handleLocalRequests() {
	http.HandleFunc("/", NewRootHandler)
	log.Printf("Started Webserver on port %d.\n", viper.GetInt("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",viper.GetInt("port")), nil))
}
func main() {
	rand.Seed(time.Now().Unix())
	err := initConfig()
	if err == nil {
		log.Printf("Initialized...")
		if viper.GetBool("RUN_LOCAL") {
			log.Printf("Will run local")
			handleLocalRequests()
		} else {
			log.Printf("Will run as Lambda, because RUN_LOCAL is not set to true")
			lambda.Start(LambdaHandler)
		}
	} else {
		log.Fatalln(err.Error())
	}
}

