package telego

import (
	"testing"
)

var Bot Load = Conf("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")

func TryConf(t *testing.T) {
	Bot := Conf("811530665:AAG5X41LQS5tJbwBaTmbs8tVXgeWYGhqYrM")
	if Bot.Token == "" {
		t.Fatalf("TOKEN NULO")
	}
}

func TestSendMessage(t *testing.T) {
	try := Bot.Send_Message(476036515, "SendMessage funcionando!")
	if try != 200 {
		t.Fatalf("Something went wrong in SendMessage.\nStatus Code:[%d]", try)
	}
}
func TestReply_TO(t *testing.T) {
	Config.Updated.MessageID = 1
	try := Bot.Reply_To(Config.Updated, 476036515, "ReplyTo funcionando")
	if try != 200 {
		t.Fatalf("Something went wrong in SendMessage.\nStatus Code:[%d]", try)
	}
}

func TestSendPhoto(t *testing.T) {
	try := Bot.SendPhoto(476036515, "AgADAQADF6gxG0nGsEbYLmuxSm6YlGwdFDAABI1ha-i8vPBpih4CAAEC")
	if try != 200 {
		t.Fatalf("Something went wrong in SendPhoto.\nStatus Code:[%d]", try)
	}
}

func TestSendAudio(t *testing.T) {
	try := Bot.SendAudio(476036515, "CQADAgADmIoDAAEcg5gKx_YQOaoQbjoC")
	if try != 200 {
		t.Fatalf("Something went wrong in SendAudio.\nStatus Code:[%d]", try)
	}
}

func TestSendDocument(t *testing.T) {
	try := Bot.SendDocument(476036515, "BQADAQADTQADScawRi4Q_TGQxqu4Ag")
	if try != 200 {
		t.Fatalf("Something went wrong in sendDocument.\nStatus Code:[%d]", try)
	}
}

func TestSendVideo(t *testing.T) {
	try := Bot.SendVideo(476036515, "BAADAQADbAADMIGwRlHrrddJsZ3KAg")
	if try != 200 {
		t.Fatalf("Something went wrong in SendVideo.\nStatus Code:[%d]", try)
	}
}

func TestSendAnimation(t *testing.T) {
	try := Bot.SendAnimation(476036515, "CgADBAADDQQAAr8ZZAeC0k40J-kcJQI")
	if try != 200 {
		t.Fatalf("Something went wrong in SendAnimation.\nStatus Code:[%d]", try)
	}
}

func TestSendVoice(t *testing.T) {
	try := Bot.SendVoice(476036515, "AwADAQADbgADMIGwRpptt4iJrzjsAg")
	if try != 200 {
		t.Fatalf("Something went wrong in SendVoice.\nStatus Code:[%d]", try)
	}
}

func TestSendLocation(t *testing.T) {
	try := Bot.SendLocation(476036515, -37.824075, -37.824075)
	if try != 200 {
		t.Fatalf("Something went wrong in SendLocation.\nStatus Code:[%d]", try)
	}
}





func TestFoward(t *testing.T) {
	Config.Updated.Chat.ID = 476036515
	Config.Updated.MessageID = 2
	try := Bot.ForwardMessage(Config.Updated, 476036515)
	if try != 200 {
		t.Fatalf("Something went wrong in FowardMessage.\nStatus Code:[%d]", try)
	}
}

func TestGetMe(t *testing.T) {
	me := Bot.GetMe()
	if !me.Ok {
		t.Fatalf("Eror in getMe!")
	}
}
