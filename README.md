# AWS Login - Simple login tool
This is a really simple cli tool designed to solve a simple problem of logging into the AWS CLI when MFA is enabled. There are other feature rich tools out there but for me i wanted something super small and lightweight and easy to understand/use.
## Installation
The tool can be built locally for all operating systems with docker by running the following command
```
make
```
A windows specific version can be built locally with docker running the following command
```
make windows
```
Alternatively download the latest release and rename to aws-login or aws-login.exe (windows).

## Usage
The tool requires a simple TOML based config file .aws-login located in the users $HOME folder. Example below
```
[aws]
mfaarn = "arn:aws:iam::<account>:mfa/<user login>"
region = "<region>"
profile = "mfa"
```
The tool can be executed by moving it into the path or running locally. This uses the config in the config file during execution
```
aws-login mfa -c <MFA Code>
```

## Alternative solutions
Other more feature rich nice implementions include 
* https://github.com/broamski/aws-mfa (python)

## Future ideas
* Use GOCUI to add a gui over the top https://github.com/jroimartin/gocui as demoed in https://github.com/jesseduffield/lazydocker







