# mbkm-logbook-translate

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GCP](https://img.shields.io/badge/Google%20Cloud%20Platform-4285f4?style=for-the-badge&logo=google-cloud&logoColor=white)
![LICENSE](https://img.shields.io/github/license/tudemaha/mbkm-logbook-translate?style=for-the-badge)

Translate MBKM Logbook for final report or other needs.

## Prerequisites 📋
- [Go](https://go.dev/) >= 1.19.0
- [Google Cloud Platform](https://cloud.google.com) with billing account connected

## How to Use ✨
### A. Preparing GCP
1. Create or use existing Google Cloud Project
2. Make sure billing account connected to that project
3. Enable [Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com) on your project
4. Create or use existing service account associated with the Cloud Translation API
5. Create and download the JSON service account key
6. Copy your Project ID to use in .env

### B. Preparing MBKM JWT Access Token
1. Login to your MBKM account
2. Open DevTool (`Ctrl + Shift + I` or `F12`)
3. Go to `Application` tab > `Storage` > `Cookies`
4. Select MBKM URL on the list
5. On the right column select `mbkmtoken`
6. Copy the JWT to use in .env

### C. Preparing `mbkm-logbook-translate`
1. Clone this repo then `cd` to that directory
```bash
git clone https://github.com/tudemaha/mbkm-logbook-translate.git
cd mbkm-logbook-translate
```
2. Move the downloaded service account to this directory
3. Fetch dependencies
```bash
go mod download
go mod tidy
```
4. Create build of `mbkm-logbook-translate`
```bash
go build -o mbkm-logbook-translate cmd/main/main.go
```
5. Prepare the .env
    - Duplicate .env.example and rename to .env
    - Fill all the required key value pairs
        - `KEY_PATH`: path of your service account key
        - `JWT`: JWT access token you've copied from MBKM cookies (Section B)
        - `PROJECT_ID`: project ID copied from Section A
        - `BASE_URL`: the base URL of MBKM logbook API (DON'T remove trailing slash)  
            `https://api.kampusmerdeka.kemdikbud.go.id/studi/report/perweek/<your_active_activity_id>/`
6. Show the `help` of `mbkm-logbook-translate` (Optional)
```bash
./mbkm-logbook-translate --help
```
7. Run the `mbkm-logbook-translate`
```bash
# using default flag's value
./mbkm-logbook-translate
# using defined weeks
./mbkm-logbook-translate -week 1,2,3
```
8. Translated loogbooks stored in `./translated` directory in JSON format

## License 📝
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.