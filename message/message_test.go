package message

import (
	"bytes"
	"github.com/unknown321/fuse/coder"
	"github.com/unknown321/fuse/tppmessage"
	"strings"
	"testing"
)

func TestMessage_Decode(t *testing.T) {
	type fields struct {
		Compress      bool
		Data          string
		OriginalSize  int
		SessionCrypto bool
		SessionKey    string
	}
	type args struct {
		src       []byte
		origSize  int
		isRequest bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "CMD_GET_URLLIST",
			fields: fields{
				Compress:      false,
				Data:          `{"lang":"ANY","msgid":"CMD_GET_URLLIST","region":"REGION_ALL","rqid":0}`,
				OriginalSize:  71,
				SessionCrypto: false,
				SessionKey:    "",
			},
			args: args{
				src:       []byte(`YnHdLj/1b4RBZXynl0xG1B0domEayp/1lLk99kX4wjo7NblxIkhv1ByeVvdNenjEJTavALlsfZfSgBpLKuCMvHkaMHdNOW9g+4ytGq/cFcXOpW6W3rjoDzBVAFXLVj+HRATx/hb68EX3+00fDqDfc0/wdXEaV+G7h5Zc4M2QoF5juLcqskL1iLDYQlLVsTH5VCgC7mK204ygBrK6BopI6RZN6pX+6R+lfT/01GExQVs=`),
				origSize:  71,
				isRequest: true,
			},
		},
		{
			name: "CMD_GET_SVRLIST",
			fields: fields{
				Compress:      false,
				Data:          `{"crypto_type":"COMMON","flowid":null,"msgid":"CMD_GET_SVRLIST","result":"NOERR","rqid":0,"server_num":0,"svrlist":[],"xuid":null}`,
				OriginalSize:  130,
				SessionCrypto: false,
				SessionKey:    "",
			},
			args: args{
				src:       []byte(`YnHdLj/1b4RBZXynl0xG1B0domEayp/1qGpbTUBIV8uyRPq87pkaswhk8kOMgKEMoiMEJi6h+1a4v5Ex0BrGdRzIKPPhe7yVa9zD8GTOmjUjj1UrxGs5dp07ax3H3wCcvH7nOnm0x39ejim7mgq8e60vEG5ZGpPlWeI8pKdLeo0gmaEGuv5lXFfSWnGcvaJU3ES9C82WwatrCtAOqyoUf+QwDriep3tPSIhUdpfmV2TLzDSGODY5u74+P8oKh9M219X6AIhPpopWpFEie9tWj4v3SS0umWYrNOzki90IR2lmE5WAfJ4klle6lutNc4Kr0vkLh9jKNm0=`),
				origSize:  130,
				isRequest: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coder.Coder{}
			err := c.WithKey(nil)
			if err != nil {
				t.Fail()
			}
			m := &Message{
				Compress:      tt.fields.Compress,
				Data:          tt.fields.Data,
				OriginalSize:  tt.fields.OriginalSize,
				SessionCrypto: tt.fields.SessionCrypto,
				SessionKey:    &tt.fields.SessionKey,
			}
			m.WithCoder(&c)
			err = m.Decode(tt.args.src)
			if err != nil {
				t.Fatalf("%s\n", err.Error())
			}

			if m.OriginalSize != tt.args.origSize {
				t.Fatalf("data size not equal, got %d, want %d", m.OriginalSize, tt.fields.OriginalSize)
			}

			if strings.Compare(m.Data, tt.fields.Data) != 0 {
				t.Fatalf("data not equal, got %s, want %s", m.Data, tt.fields.Data)
			}

			if m.Compress != tt.fields.Compress {
				t.Fatalf("compress mismatch, want %t", tt.fields.Compress)
			}
		})
	}
}

func TestMessage_Encode(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		src     Message
		wantErr bool
	}{
		{
			name: "CMD_GET_URLLIST",
			src: Message{
				Compress:      false,
				MData:         []byte("{test}"),
				SessionCrypto: false,
				MsgID:         tppmessage.CMD_GET_URLLIST,
			},
			want: []byte(`YnHdLj/1b4RBZXynl0xG1B0domEayp/1lKyOrmgtL/2InpcwEjMLC9bkCAwD5szyA7dp9KaZfNuw2EJS1bEx+VQoAu5ittOMoAayugaKSOkWTeqV/ukfpUEl05FvqCQt`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &tt.src
			m.OriginalSize = len(m.MData)

			c := coder.Coder{}
			err := c.WithKey(nil)
			if err != nil {
				t.Fail()
			}

			m.WithCoder(&c)

			//slog.SetLogLoggerLevel(slog.LevelDebug)

			got, err := m.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %s, wantErr %v", err.Error(), tt.wantErr)
				return
			}

			if bytes.Compare(got, tt.want) != 0 {
				t.Errorf("Encode() got = %s, want %s", got, tt.want)
			}
		})
	}
}
