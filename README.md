## Simple Git wrapper to make it easier for me to work with Git

### Installation

```bash
curl --fail --location --output gclg https://github.com/frixaco/gclg/releases/download/0.0.1/gclg
sudo chmod +x gclg
sudo mv gclg /usr/local/bin/gclg
```

### Features

- To clone a repo: `gclg <user>/<repo>` (instead of `git clone git@github.com:<user>/<repo>.git`)

### Supported platforms

- Linux (tested only on Arch as of now)