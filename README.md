go build -ldflags="-H windowsgui"

fyne package -os windows -icon ./theme/icon.png -name TinsRPC
fyne package -os darwin -icon ./theme/icon.png -name TinsRPC
fyne package -os linux -icon ./theme/icon.png -name TinsRPC
