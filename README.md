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
```