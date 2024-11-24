package gentoken

import (
	"context"
	"fmt"
	"reflect"

	SmartApi "github.com/angel-one/smartapigo"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/date"
	mongo "github.com/slahoty91/tradingBot/copyOrderAngleOne/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AngleOneStru struct {
	APIKey        string `bson:"apiKey"`
	ClientID      string `bson:"clientID"`
	Pass          string `bson:"pass"`
	LastLoginDate string `bson:"lastLoginDate"`
	AccToken      string `bson:"accToken"`
}
type User struct {
	ID         string       `bson:"_id"` // Needs to be string for bson.ObjectId
	APIKey     string       `bson:"apikey"`
	APISecrete string       `bson:"apiSecrete"`
	AngleOne   AngleOneStru `bson:"angleOne"`
}

var ABClient *SmartApi.Client

func GetClient() (User, *SmartApi.Client) {
	filter := bson.M{"name": "SIDDHARTH LAHOTY"}
	opts := options.FindOne()
	opts.SetProjection(bson.M{"angleOne": 1})
	userCol := mongo.GetCollection("algoTrading", "userDetails")

	userData := userCol.FindOne(context.Background(), filter)
	var doc User
	if err := userData.Decode(&doc); err != nil {
		fmt.Println(err)
	}
	// fmt.Println("DOC====>>>>>", doc)
	angleOne := doc.AngleOne
	// angleOne = doc.AngleOne
	apiKey := angleOne.APIKey
	user_id := angleOne.ClientID
	pass := angleOne.Pass
	// fmt.Println(user_id, pass, apiKey, "HIIIIIIIIIII")
	ABClient = SmartApi.New(user_id, pass, apiKey)
	fmt.Println("After ABC CLIENT", ABClient)
	return doc, ABClient
}
func GenerateToken() {

	date := date.CurrentDate()
	// lastLoginDate := ""
	// filter := bson.M{"name": "SIDDHARTH LAHOTY"}
	// opts := options.FindOne()
	// opts.SetProjection(bson.M{"angleOne": 1})
	// userCol := mongo.GetCollection("algoTrading", "userDetails")

	// userData := userCol.FindOne(context.Background(), filter)
	// var doc User
	// if err := userData.Decode(&doc); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("DOC====>>>>>", doc)
	doc, ang := GetClient()
	fmt.Println("DOC from GenerateToken====>>>>>", doc, ang)
	lastLoginDate := doc.AngleOne.LastLoginDate
	angleOne := doc.AngleOne
	fmt.Println(angleOne, "<======AngleOne")
	fmt.Println(date, lastLoginDate, reflect.TypeOf(date), reflect.TypeOf(lastLoginDate), date == lastLoginDate)
	if date == lastLoginDate {
		fmt.Println("Equal date, session in progress")
		var newSession string
		fmt.Println("Generate New session regardless, press Y to generate, N to exit program :-")
		fmt.Scanln(&newSession)
		var optionsTrue = [7]string{"Y", "y", "yes", "Yes", "YEs", "yES", "yEs"}
		var optionsFalse = [7]string{"N", "n", "No", "no", "nO", "n0", "N0"}
		isTrue := false
		isFalse := false
		for _, option := range optionsTrue {
			if newSession == option {
				isTrue = true
				genAccTok()
				break
			}

		}
		for _, option := range optionsFalse {
			if newSession == option {
				isFalse = true
			}
		}
		if !isTrue && !isFalse {
			GenerateToken()
		}

	} else {
		fmt.Println("Date not equal")

		genAccTok()
	}
}

func genAccTok() {

	date := date.CurrentDate()
	// apiKey := angleOne.APIKey
	// user_id := angleOne.ClientID
	// pass := angleOne.Pass
	// ABClient := SmartApi.New(user_id, pass, apiKey)
	// _, ABClient := GetClient()
	var totp string
	fmt.Println("Enter TOTP :-")
	fmt.Scanln(&totp)
	fmt.Println(totp, "<====TOTPPPP")

	fmt.Println("Client :- ", ABClient)

	// User Login and Generate User Session
	session, err := ABClient.GenerateSession(totp)

	if err != nil {
		fmt.Println(err)
		return
	}
	session.UserSessionTokens, err = ABClient.RenewAccessToken(session.RefreshToken)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User Session Tokens :- ", session.UserSessionTokens)
	session.UserProfile, err = ABClient.GetUserProfile()

	if err != nil {
		fmt.Println(err.Error(), "ERRORRRRRRRRR")
		return
	}

	fmt.Println("User Profile :- ", session.UserProfile)
	fmt.Println("User Session Object :- ", session.AccessToken, session.LastLoginTime)
	userCollec := mongo.GetCollection("algoTrading", "userDetails")
	filter := bson.M{
		"name": "SIDDHARTH LAHOTY",
	}
	update := bson.M{

		"$set": bson.M{
			"angleOne.lastLoginDate":         date,
			"angleOne.accToken":              session.AccessToken,
			"angleOne.feedToken":             session.FeedToken,
			"angleOne.refreshToken":          session.RefreshToken,
			"angleOne.lastLoginTimeAngleOne": session.LastLoginTime,
		},
	}
	updateResult, err := userCollec.UpdateOne(context.Background(), filter, update)

	fmt.Println(updateResult, err)
}
