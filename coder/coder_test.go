package coder

import (
	"reflect"
	"strings"
	"testing"
)

type message struct {
	Encoded string
	Decoded string
}

func TestCoder_Decode(t *testing.T) {
	type fields struct {
		Message message
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "CMD_GET_URLLIST", fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1lLk99kX4wjo7NblxIkhv1ByeVvdNenjEJTavALlsfZfSgBpLKuCMvHkaMHdNOW9g+4ytGq/cFcXOpW6W3rjoDzBVAFXLVj+HRATx/hb68EX3+00fDqDfc0/wdXEaV+G7h5Zc4M2QoF5juLcqskL1iLDYQlLVsTH5VCgC7mK204ygBrK6BopI6RZN6pX+6R+lfT/01GExQVs=`,
					Decoded: `{"compress":false,"data":"{\"lang\":\"ANY\",\"msgid\":\"CMD_GET_URLLIST\",\"region\":\"REGION_ALL\",\"rqid\":0}","original_size":71,"session_crypto":false,"session_key":""}`,
				},
			},
		},
		{
			name: "CMD_GET_URLLIST response, compressed",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4SBvSa1/0bYhcd4UAB70VUxUp9R+E7nlQqbd/CfasI2cHERrSHJdSMpyXeXDjDyVZA9ZW+a9XxS/njMyTWS86ztRmyJ6yr6RICSRqkSq/14sMwyaWtw5heV+Sz+EYH1y2pbbsdaoU5hYHVAPWVjLqPUc8dYBIG7yZXPEmu0H3/N2z6uhACqfTgyrJvd8dOR+t4ekBWWFsLFTzO4yUOlMIk4kV7NHmMm2QPwfUE8/YY05Pog0Cu6r6wGSIZJghUfd8Ko9jzdAqyHkq/07vhNMcet3RBbi7BJty7zbN01+7nuywzb3bZh3MnWjqhs3swuQUADCo4k4+0PrfmWJY5kTwNXGALVleozNL2CR0pLVHimlfKohUvBlLP6kOzrli7EIC5EBqPPUTTptMmfkE2PTm9wDf1Lk+MPib7YpYqgOsPbIjONlpiur+19VdzkXTpOJBir1mVkNGGxdgTntsNuRyLpE1VkE7ngGSnX6jNAs8vFF65LERPnvE6jnsAlzQt7YDRpzTvu0nOD4I5YAzAdBaJkaiPelF7/nmdTt5T1ic2LNbWcg37hPPar2PTXrReJtq6xwMvNY/IzZMamJefkWTTC`,
					Decoded: `{"compress":true,"data":"eNqd0l1LwzAUBuD/cq5jWxAv7F3twhy0q9ROL0RK3GIXTNKYj80x+t9Np5OKdMhuAidv8hw4yR6W\r\neqdsW9udohBDWuR5MQcEr7zdshXE0nGOQJimLyDNJ/UUV/WizLLZfeXPaWoctz6aF7gs+433/mSE\r\nwGlec2Z89rSHb36aVBgOkS/W1ioTh6FojFXqoiGCBm+tJIK1kjNJg2UrQp8YK8KGWOovbqg2PoX4\r\nukM/6CO+OcsUhMkxEy+y5Bf6P3NLX0LqODkswYbooX858G9xUuXJ3UgLotSpDmtKrCBqaEcDe4If\r\nZunomE/YfyYSdc9f7yidgPgKwYc7/onuEwCNwqk=\r\n","original_size":570,"session_crypto":false,"session_key":null}`,
				},
			},
		},
		{
			name: "CMD_GET_SVRLIST",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1lLk99kX4wjo7NblxIkhv1ByeVvdNenjEJTavALlsfZfSgBpLKuCMvItPlQX75xlZ0zvQUV2ISsgKgvkURGBV62nNO+7Sc4PgjG0XFUSpOd2jrjLz0D+aIJBC6VS6fT6TFs+cVW03735Cjm5lNz8sSyFMDh0PzigW5J5kmMhKqKc=`,
					Decoded: `{"compress":false,"data":"{\"lang\":\"ANY\",\"msgid\":\"CMD_GET_SVRLIST\",\"rqid\":0}","original_size":49,"session_crypto":false,"session_key":""}`,
				},
			},
		},
		{
			name: "CMD_AUTH_STEAMTICKET",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1h6yLK83LXI4zDktCs7Vs0Mhm715aaBJXJs7vkIRLGroxCTP9sTY8wmvcw/Bkzpo1JYLn5AkP2dh371LsZ/MbjtLKOtSKp5oE4lI5G6240U380+NRPXDllh8BDgbzscpgjJfoou+Qvoxe5M+l0nntVpLFiFex/WBLw9Z5AyKJLVC5jYMQW68kcjljq1jhYyhB+h1f3SKdqWjYJYGTaqMRSVepTj8lCqGPtLgSHTW1Fubv8wvAeMxtJVFf2XoRfPSXQscNtEMCSIMp6jAOpuBEKESajdI/LR8JlyW7yP7U+eVx0VO8wQY6IFIYxEU4CXsHn6QJWeGPSgquVO2FHCSxVic5zzVOrtpGNC3cgXClmtfCLqafz95KureP49/cpwzGktfC9bu+AFRUfS+u8mzunsBqMkGYlco+fZ5qauyzVhXUJzz1Zs2nXo0LXfoIjcRJxGsIuQLFqqZQA3IzIJr1RY1DceywSRpdch+NiqTf7lsdITTN5vwyz2AJyIAYYAWrhdZYhLepMh/8fY5T3zCSY/7G850BrJ475qbnMaGLA7ps5yv6Yd2J+e1AweL/9IZiTA5hRuxYbKjlm+tw7L45TL05QpW/AUWzQv00DFrHDd9PV5FiC2nLriKm7HEPZcSjZIDqKrh4PqWb4Az12lSiMCdFqZ0MxfLkrME60F8Rf7W3mqDMU430mULHDbRDAkiDIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSvqG2CmkTuPMhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZLKICLc3eBhzbPVKBfzYgFRIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kgJ+bXaWvYJz+Li2aAYBW0ghdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZJCxw20QwJIgyF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kr6htgppE7jzIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSyiAi3N3gYc2z1SgX82IBUSF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZICfm12lr2Cc/i4tmgGAVtIIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSQscNtEMCSIMhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZK+obYKaRO48yF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1ksogItzd4GHNs9UoF/NiAVEhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSAn5tdpa9gnP4uLZoBgFbSCF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kkLHDbRDAkiDIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSIXSHwAGYtZIhdIfAAZi1kiF0h8ABmLWSXLYojWn1fijAu6DePaOkcLg5lW5RTeIidk+zCs2JK+9P8HVxGlfhu4eWXODNkKBe6h9S0rULkCaBZnIwQBzhjDKzUwL7fVQN3dyGJ9iUwezy1GJiL0GYIE3MPv+biv8l`,
					Decoded: `{"compress":false,"data":"{\"country\":\"ww\",\"lang\":\"en\",\"msgid\":\"CMD_AUTH_STEAMTICKET\",\"region\":4,\"rqid\":0,\"steam_ticket\":\"FAAAACSH9hD8injFdEfUAQEAEAFJmKBYGAAAAAEAAAACAAAAnhCZJQMBqMDvP60BAQAAAOIAAABi\\r\\nAAAABAAAAHRH1AEBABAB1GMEAJ4QmSVlCqjAAAAAAHgHmlj4trVYAQDVIgEACACQAAYAAAAMNAYA\\r\\nAAAgNAYAAAA0NAYAAAA+NAYAAABINAYAAABSNAYAAABcNAYAAAAAAAs4HyYawguSD06gown98wvJ\\r\\nHdOMhU594WvucMrpS5MKoVjfkia+6MRY1dpytBAClqZUtpkoV3Cy0wZ5rvxvCuEvdh9ruj39aRCe\\r\\nkUaGHSTQtREbEY1WRS9y8QfgZNnHqzh5+ZxQoCd3wd2Wm37FXPG9Y+xgph/Iy2ifAYz6BZloAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\\r\\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"steam_ticket_size\":282}","original_size":1557,"session_crypto":false,"session_key":""}`,
				},
			},
		},
		{
			name: "6, CMD_GET_SVRTIME",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1S25XF7R+9x/tCTuJ8y8ehC028BLN6hpWb1BKhkvZEq1Z4jykp0t6jX6eDMnMijgJszV2RzwHLYBr44DXQhZC6UKObmU3PyxL/vGm9TScTopepBXbFf1bklakUSJ721aPXU/zis2cTZoElxq7cHOH/A==`,
					Decoded: `{"compress":false,"data":"{\"msgid\":\"CMD_GET_SVRTIME\",\"rqid\":0}","original_size":36,"session_crypto":false,"session_key":""}`,
				},
			},
		},
		{
			name: "7, CMD_AUTH_STEAMTICKET",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1qX4HSxZr8MwvezTSGdbUTw9gdwfkr4j9lDhlPHTg9EKoND6XIJJSo5tyxg++gfjXXj9VM/hC+M1iMvlcjGeF97kEDatZFP9aqBpybjbR2wTVTIgdxm+O7qmyVDBzqW0kdbaD0f15tgbilFgW+sQlB44Hpc8vsDxlPmm9uVwDxbVM1uIWd3Vi/lhsniFguXyZedmR8FGMk7/fqB5WBsTeBhsLWg0vO97JE1PbGpGe75sxtBBnOqkS8zC3dbRJwDakrLQPIhmjpnU+MeeKV3GDJZCzoGZ4Xwa+0zvQUV2ISsjcRL0LzZbBqzSFhUtTwwt6Ti7BZHE6dlUdZBoV3LkCzmTlBj3B2P3byplLUwTylNcVSezKuXzK8xcXyLGHn6pA5hYVKhzBUEuugmUmAXyHFCUugIg+QuW7HJ05BWWisxOZycIDScu2H7NY1/I0Gf1RSM4fnKfS9iRl1EVxRtROHiOOCp9s8loinM9kZUkt3BsDMCp2FNlkXIielzASMwsL1uQIDAPmzPJIMGuCHv4dwIFmcjBAHOGMMrNTAvt9VA3d3IYn2JTB7PLUYmIvQZgg3TzhCQZFl+Hz9g6Y1Bl6KQ==`,
					Decoded: `{"compress":false,"data":"{\"account_id\":\"76561197960287950\",\"crypto_type\":\"COMMON\",\"currency\":\"NOK\",\"flowid\":null,\"loginid_password\":\"0877a5c54790d7051c877ec648c9b07f\",\"msgid\":\"CMD_AUTH_STEAMTICKET\",\"result\":\"NOERR\",\"rqid\":0,\"smart_device_id\":\"anFvNGA5ZlRLU1ZDN1FqeUNLZ25cNmZrOHpcQXhXQnBXU5BVUm9nSzNcNnpjQ2p2dDQ5dWd5R1NNeVZlZU50ZGZ4V2tJMEc5YUhZSHVLdmM=\",\"xuid\":null}","original_size":540,"session_crypto":false,"session_key":null}`,
				},
			},
		},
		{
			name: "8, CMD_GET_SVRTIME_resp",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1qGpbTUBIV8uyRPq87pkaswhk8kOMgKEMoiMEJi6h+1ZWo5ZVN1fhX+anqoCtXaP8mmSLBuKkh+VIiFR2l+ZXZDuTlEwlGnXXE1PbGpGe75sDuG8rPPiZ6iJoKAr6LdTNW5sVYf1qsAUvnswOq2GbHjjriN3TiPwGpWVfrrCLUvJD1WqUmzNjj2Lkttvnel9WBmiQhGw4hSGeHY92rqIGow8gAw2BMFSwsNhCUtWxMflUKALuYrbTjKAGsroGikjpFk3qlf7pH6VBJdORb6gkLQ==`,
					Decoded: `{"compress":false,"data":"{\"crypto_type\":\"COMMON\",\"date\":1486920003,\"flowid\":null,\"msgid\":\"CMD_GET_SVRTIME\",\"result\":\"NOERR\",\"rqid\":0,\"xuid\":null}","original_size":120,"session_crypto":false,"session_key":null}`,
				},
			},
		},
		{
			name: "9, CMD_REQAUTH_HTTPS",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1oOB/aaqGUJ/xPc/5EZDE0la0zxbXormATaj+W9q/T3w3ckAPxuEUnc2YoiYzmtgCldfuFRlpuBpr3MPwZM6aNZwokmYO1pZq+pSM3eiOW08fIDMKat5KsyvWS/5lX30FQBCV5eWXBukvBdnFXbd1i6VlX66wi1LykK3VF4ZikuBBHP61navffgnHjYxoi7TL/sJFjGZBMAO9g21isiWfyIXlA2dVHjb4HGbu8Y8HFCUplkU6T8Js3XooI83yOmupT/B1cRpX4buHllzgzZCgXlzmK/QcpvKh/HrjydR2cNoRYe0d61cH+KmmC5WWwiGrB7crpfDoWQd9dDQnvxICZQ==`,
					Decoded: `{"compress":false,"data":"{\"hash\":\"zpiPLHJUKxC3u3siy9NLhg==\",\"is_tpp\":1,\"msgid\":\"CMD_REQAUTH_HTTPS\",\"platform\":\"Steam\",\"rqid\":0,\"ugc\":1,\"user_name\":\"76561197960287930\",\"ver\":\"NotImplement\"}","original_size":163,"session_crypto":false,"session_key":""}`,
				},
			},
		},
		{
			name: "10, CMD_REQAUTH_HTTPS_resp",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/16jFY43jA027f/KuDsCgyolGZH6o3LqSSeE7Lo2uVhg9nfK3R/mriCAEZRI59nlP2PqvlQ+T/VF6pUDHeXJaCyzydOGl3qeUmSq5Nql0Z7KmbcsYPvoH4114/VTP4QvjNYjL5XIxnhfc9/LpysET6tc3Qe/BkTvgomRHCbCFdTnWc8FAGJp8bN7ko9N7Xorbs+wWNtBIjlDCtjg1KUASopOOEs8FQZjpz6K4Xx1fH+uzCLx6tL3Ekt9BxZGOK0Q/wweFqtBGXvKfatrLE5PFXQyU2rwC5bH2XpehR/r7XrfT3XhMM6prPN2kksw2Y4rhDXo4pu5oKvHutLxBuWRqT5VniPKSnS3qNIJmhBrr+ZVz/yWzzFDYcTlv4JmmlVY4FPi2isvs8PPKQ/3m3nh8C+6HC1kgBh0v61q0PDoyM94M0hYVLU8MLek4uwWRxOnZVHWQaFdy5As5k5QY9wdj928qZS1ME8pTXFUnsyrl8yvMXF8ixh5+qQOYWFSocwVBLroJlJgF8hxQlLoCIPkLluxydOQVlorMTmcnCA0nLth+zWNfyNBn9UUjOH5yn0vYkZdRFcUbUTh4jjgqfbPJaIlFjMyxu1914zZQcOSFnDdciGC6HhMx+L4R6fp/pGw71CoSJk1PRE9dD1WqUmzNjj2Lkttvnel9WBmiQhGw4hSGeHY92rqIGo7+7lRagi0RtsNhCUtWxMflUKALuYrbTjKAGsroGikjpFk3qlf7pH6VBJdORb6gkLQ==`,
					Decoded: `{"compress":false,"data":"{\"aes_key\":null,\"cbc_iv\":null,\"crypto_key\":\"t1AfMMWEWHb2lmv+P2lOyA==\",\"crypto_type\":\"COMMON\",\"flowid\":null,\"heartbeat_sec\":60,\"hmac_key\":null,\"inquiry_id\":765071051,\"is_use_apr\":0,\"msgid\":\"CMD_REQAUTH_HTTPS\",\"result\":\"NOERR\",\"rqid\":0,\"session\":\"88e282c21a494722a3867b2d945faecb\",\"smart_device_id\":\"anFvNGA5ZlRLU1ZDN1FqeUNLZ25cNmZrOHpcQXhXQnBXU5BVUm9nSzNcNnpjQ2p2dDQ5dWd5R1NNeVZlZU50ZGZ4V2tJMEc5YUhZSHVLdmM=\",\"timeout_sec\":200,\"user_id\":12345,\"xuid\":null}","original_size":454,"session_crypto":false,"session_key":null}`,
				},
			},
		},
		{
			name: "11, encrypted",
			fields: fields{
				Message: message{
					Encoded: `YnHdLj/1b4RBZXynl0xG1B0domEayp/1wMR0SOcV9/LioSSJDA2uIP2IuhKHdoixAjm5Jy7BQu2heWhqNaNvkL5iP8TUOosEJYzskIIu79OzENC9iO73b6Xw1Y0wtvoxnwcMIlQewbbv8yWABfFkIqbClrNIZYgQZ5LV3OLqp/FZqCtVESmKBoj+BaZe5ZFriWnzTF3pGEpXMcKEirF0p3At8LJ4kBuiX8UtYnujKxCBKLXXYKA6EzMtUn2PFqlDZIsoZdNwOqfdiCsJmJKSohPARtJhvYb5xnmqDWuhj07Oj1MkkgGBBT7Kh5NRHdu7ggdGLIY9c3vnZRQGwhbh6X2tk2LNHM2Y836VdT202oP/t1tG9Y5BrjmPQDNpqA2F8sfGBHf9x9fEm9+nmb/f+RYHw31s6SpWpyzmnVziDzPyx8YEd/3H18Sb36eZv9/5EWlvZmbQNI01Sk78M78AU/AMOjUWiASKziQ5sbB2EHq+Pj/KCofTNkZoUZte9E4SVqRRInvbVo+L90ktLplmK35rTUfvsnLk7TSdQf/oWw2j18E29pPt8drw7cywwTt+AWJpKpyTK/V+s0iN9hWd01v2PiDnxb8p`,
					Decoded: `{"compress":false,"data":"5z2zpfCvk721C1wO9wxpyEc1LX9U2OLekoWIHm9RVH085Y+2G5ZqGFU3Qm9jrvD45oDfKfJ0dR3s\r\nds2WlUSRVrk0QZuMiKovgWuqbI2zswSLbeHNnPWkBhcdpmSma7z0f6PKHy63RZxtnDsdm9boDfhe\r\nHrkkz5iQYbql+XLarZxdgPvS1zS+LuS8r4c8kSRddrqRX41CiQNvmQW3VfjaCzaJxkL2STljtpJn\r\nC4LXJ5M7WStId6aWPG9eYPPVf2HBb15g89V/YcFvXmDz1X9hwW9eYPPVf2HBb15g89V/YcFvXmDz\r\n1X9hwQC8CGPo+g0Y","original_size":234,"session_crypto":true,"session_key":"88e282c21a494722a3867b2d945faecb"}`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coder{}
			_ = c.WithKey(nil)
			res := c.Decode([]byte(tt.fields.Message.Encoded))
			if strings.Compare(string(res), tt.fields.Message.Decoded) != 0 {
				t.Fatalf("have\n%s\n, want\n%s\n", res, tt.fields.Message.Decoded)
			}
		})
	}
}

func TestCoder_Encode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "CMD_GET_URLLIST req",
			args: args{
				data: []byte(`{"compress":false,"data":"{\"lang\":\"ANY\",\"msgid\":\"CMD_GET_URLLIST\",\"region\":\"REGION_ALL\",\"rqid\":0}","original_size":71,"session_crypto":false,"session_key":""}`),
			},
			want: []byte(`YnHdLj/1b4RBZXynl0xG1B0domEayp/1lLk99kX4wjo7NblxIkhv1ByeVvdNenjEJTavALlsfZfSgBpLKuCMvHkaMHdNOW9g+4ytGq/cFcXOpW6W3rjoDzBVAFXLVj+HRATx/hb68EX3+00fDqDfc0/wdXEaV+G7h5Zc4M2QoF5juLcqskL1iLDYQlLVsTH5VCgC7mK204ygBrK6BopI6RZN6pX+6R+lfT/01GExQVs=`),
		},
		{
			name: "CMD_GET_URLLIST resp",
			args: args{
				data: []byte(`{"compress":true,"data":"eNqd0l1LwzAUBuD/cq5jWxAv7F3twhy0q9ROL0RK3GIXTNKYj80x+t9Np5OKdMhuAidv8hw4yR6W\r\neqdsW9udohBDWuR5MQcEr7zdshXE0nGOQJimLyDNJ/UUV/WizLLZfeXPaWoctz6aF7gs+433/mSE\r\nwGlec2Z89rSHb36aVBgOkS/W1ioTh6FojFXqoiGCBm+tJIK1kjNJg2UrQp8YK8KGWOovbqg2PoX4\r\nukM/6CO+OcsUhMkxEy+y5Bf6P3NLX0LqODkswYbooX858G9xUuXJ3UgLotSpDmtKrCBqaEcDe4If\r\nZunomE/YfyYSdc9f7yidgPgKwYc7/onuEwCNwqk=\r\n","original_size":570,"session_crypto":false,"session_key":null}`),
			},
			want: []byte(`YnHdLj/1b4SBvSa1/0bYhcd4UAB70VUxUp9R+E7nlQqbd/CfasI2cHERrSHJdSMpyXeXDjDyVZA9ZW+a9XxS/njMyTWS86ztRmyJ6yr6RICSRqkSq/14sMwyaWtw5heV+Sz+EYH1y2pbbsdaoU5hYHVAPWVjLqPUc8dYBIG7yZXPEmu0H3/N2z6uhACqfTgyrJvd8dOR+t4ekBWWFsLFTzO4yUOlMIk4kV7NHmMm2QPwfUE8/YY05Pog0Cu6r6wGSIZJghUfd8Ko9jzdAqyHkq/07vhNMcet3RBbi7BJty7zbN01+7nuywzb3bZh3MnWjqhs3swuQUADCo4k4+0PrfmWJY5kTwNXGALVleozNL2CR0pLVHimlfKohUvBlLP6kOzrli7EIC5EBqPPUTTptMmfkE2PTm9wDf1Lk+MPib7YpYqgOsPbIjONlpiur+19VdzkXTpOJBir1mVkNGGxdgTntsNuRyLpE1VkE7ngGSnX6jNAs8vFF65LERPnvE6jnsAlzQt7YDRpzTvu0nOD4I5YAzAdBaJkaiPelF7/nmdTt5T1ic2LNbWcg37hPPar2PTXrReJtq6xwMvNY/IzZMamJefkWTTC`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coder{}

			_ = c.WithKey(nil)
			if got := c.Encode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("have\n%s\n, want\n%s\n", got, tt.want)
			}
		})
	}
}
