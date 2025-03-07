# imazes
Generate images from your terminal for free

# Usage
From source
```bash
git clone https://github.com/rohitaryal/imazes
cd imazes
go build ./cmd/imazescli/main.go
```

Or get pre-compiled binary from [releases](https://github.com/rohitaryal/imazes/releases)

# Available Options

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