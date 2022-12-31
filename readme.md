# TFGo

Analyze terraform (*.tf) scripts to find insights from them

## Installation

### From Github releases page

Go to [Release page](https://github.com/danny270793/TFGo/releases) then download the binary which fits your environment

### From terminal

Get the last versión available on github

```bash
LAST_VERSION=$(curl https://api.github.com/repos/danny270793/TFGo/releases/latest | grep tag_name | cut -d '"' -f 4)
```

Download the last version directly to the binaries folder

For Linux (linux):

```bash
curl -L https://github.com/danny270793/TFGo/releases/download/${LAST_VERSION}/TFGo_${LAST_VERSION}_linux_amd64.tar.gz -o ./tfgo.tar.gz
```

For MacOS (darwin):

```bash
curl -L https://github.com/danny270793/TFGo/releases/download/${LAST_VERSION}/TFGo_${LAST_VERSION}_darwin_amd64.tar.gz -o ./tfgo.tar.gz
```

Untar the downloaded file

```bash
tar -xvf ./tfgo.tar.gz
```

Then copy the binary to the binaries folder

```bash
sudo cp ./TFGo /usr/local/bin/tfgo
```

Make it executable the binary

```bash
sudo chmod +x /usr/local/bin/tfgo
```

```bash
tfgo --version
```

## Ussage

Run the binary and pass the path to the folder where you want to look for missing variables

```bash
tfgo variables missing -path ./path/to/modules/folder
```

Or search when the variables declared are ussed

```bash
tfgo variables ussage -path ./path/to/modules/folder
```

## Follow me

* [Facebook](https://www.facebook.com/danny.vaca.9655)
* [Instagram](https://www.instagram.com/danny27071993/)
* [Youtube](https://www.youtube.com/channel/UC5MAQWU2s2VESTXaUo-ysgg)
* [Github](https://www.github.com/danny270793/)
* [LinkedIn](https://www.linkedin.com/in/danny270793)

## LICENSE

Licensed under the [MIT](license.md) License.

## Author

[@danny270793](https://github.com/danny270793)
