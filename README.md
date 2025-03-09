# 📷 imazes
Generate images from your terminal for free

## 🎨 Usage
From source
```bash
git clone https://github.com/rohitaryal/imazes
cd imazes
go build ./cmd/imazescli/main.go
```

Or get pre-compiled binary from [releases](https://github.com/rohitaryal/imazes/releases)

## 🤓 Available Options

```
imazen:
	Generate AI images from your terminal

Usage: imazen [OPTION...]

  -p,  --prompt	 Prompt string that describes an image
  -n,  --negative	 Avoid generating specific elements or characteristics
  -s,  --style	 Name of unique visual characteristics and aesthetic
  -c,  --count	 Number of images you want to generate
  -k,  --steps	 Number of iterative denoising processes the model performs
  -r,  --ratio	 Aspect ratio or size of generated image
  -h,  --help	 Generate this help

Provided By: @rohitaryal

[!] Available Styles include:

Medieval            
Vincent Van Gogh    F Dev               Low Poly            Dreamshaper-xl      Anima-pencil-xl     
Biomech             Trash Polka         No Style            Cheyenne-xl         Chicano             
Embroidery tattoo   Red and Black       Fantasy Art         Watercolor          Dotwork             
Old school colored  Realistic tattoo    Japanese_2          Realistic-stock-xl  F Pro               
RevAnimated         Katayama-mix-xl     SDXL L              Cor-epica-xl        Anime tattoo        
New School          Death metal         Old School          Juggernaut-xl       Photographic        
SDXL 1.0            Graffiti            Mini tattoo         Surrealism          Neo-traditional     
On limbs black      Yamers-realistic-xl Pony-xl             Playground-xl       Anything-xl         
Flame design        Kawaii              Cinematic Art       Professional        Flux                
Black Ink           


[!] Available ratio include:

1:1  
2:3  3:2  3:4  4:3  9:16 
16:9 9:21 21:9 
```

## 🤔 How it works?
Idea is simple, and I was able to capture these HTTP requests from [this](https://play.google.com/store/apps/details?id=ai.generated.art.maker.image.picture.photo.generator.painting) app

### Step 1: Generate a token
```http
POST /identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyB3-71wG0fIt0shj0ee4fvx1shcjJHGrrQ HTTP/2
Host: www.googleapis.com
Content-Type: application/json
X-Firebase-Client: H4sIAAAAAAAAAKtWykhNLCpJSk0sKVayio7VUSpLLSrOzM9TslIyUqoFAFyivEQfAAAA
Content-Length: 38

{"clientType":"CLIENT_TYPE_ANDROID"}
```

You will get a response in the following format, please note down the value of `idToken`:
```JSON
{
  "kind": "...",
  "idToken": "__TOKEN__WILL__BE__HERE__",
  "refreshToken": "...",
  "expiresIn": "3600",
  "localId": "..."
}
```

### Step 2: Generate an image
```http
POST /api/v1/text2image HTTP/2
Host: img-gen-prod.ai-arta.com
Authorization: __TOKEN__WILL__BE__HERE__
Content-Type: multipart/form-data; boundary=fe715f8a-7532-4c89-a461-2a193f92c487

--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="prompt"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 133

A poor developer
--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="negative_prompt"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 0


--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="style"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 6

SDXL 1.0
--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="images_num"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 1

4
--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="cfg_scale"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 1

7
--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="steps"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 2

30
--fe715f8a-7532-4c89-a461-2a193f92c487
Content-Disposition: form-data; name="aspect_ratio"
Content-Transfer-Encoding: binary
Content-Type: text/plain; charset=utf-8
Content-Length: 3

1:1
--fe715f8a-7532-4c89-a461-2a193f92c487--
```

The response will be in following format, and record_id represents our specific image generation request which we will be using to track the current status of generation status. Please note down the `record_id`
```JSON
{
  "record_id":"__RECORD__ID__HERE__",
  "status":"IN_QUEUE",
  "response":null,
  "error_code":null,
  "error_details":null,
  "seed":633479648
}
```

### Step 3: Check the status

You will get replied as this at first request:

```JSON
{
  "record_id":"__RECORD__ID__HERE__",
  "status":"IN_PROGRESS",
  "response":null,
  "error_code":null,
  "error_details":null,
  "seed":3965768918
}
```

And after few more times you will get this response for same request

```JSON
{
    "record_id": "__RECORD__ID__HERE__",
    "status": "DONE",
    "response": [
        {
            "name": "__IMAGE__NAME__HERE__",
            "url": "__IMAGE__LINK__HERE__",
            "isBlur": false,
            "MIME": "image/png"
        },
        {
            "name": "__IMAGE__NAME__HERE__",
            "url": "__IMAGE__LINK__HERE__",
            "isBlur": false,
            "MIME": "image/png"
        },
        {
            "name": "__IMAGE__NAME__HERE__",
            "url": "__IMAGE__LINK__HERE__",
            "isBlur": false,
            "MIME": "image/png"
        },
        {
            "name": "__IMAGE__NAME__HERE__",
            "url": "__IMAGE__LINK__HERE__",
            "isBlur": false,
            "MIME": "image/png"
        }
    ],
    "error_code": null,
    "error_details": null,
    "seed": 3965768918
}
```