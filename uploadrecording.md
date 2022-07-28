# Readme

## Setup

Step1] 
Open console.cloud.google.com
Create a project and enable Youtube Data API from the Library.
Fill your app information on the OAuth Consent Screen and add `.../auth/youtube.upload` as a scope.
Now under the credentials section create OAuth Client ID, Application type -> Desktop App.
Download the JSON file and save it as `client_secret.json` in root directory of this project.


Step2] 
uploadrecording.go takes in 3 parameters: file, title, and description.
Replace the `test.mov` file with your own video file you want to upload.

Now run the uploadrecording.go file with this command:
``` go run uploadrecording.go --file="test.mov" --title="Test video" --description="golang test"```


Step3]
Prompt:
```Go to the following link in your browser then type the authorization code below: 
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=*******
 Enter your code here
```
Accept all the permissions of the Google authorization box.
The `code` you provide above from your localhost is an unique authorization code which is required to create an access token so that you can talk with the API. This access token will be saved to your local machine's root directory.

```Saving credential file to: ~/.credentials/uploadrecording.json
The video with id 'YFDok-Ue_zA' is successfully initiated for upload.
```
