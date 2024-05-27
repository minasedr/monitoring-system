# Project Name: PRTG Mini Probe API Integration

## Description:

This project integrates with the PRTG (Paessler Router Traffic Grapher) Mini Probe API to collect data from an external API and send it to a PRTG monitoring system. The application periodically queries the external API, processes the data, the response time, and sends it to PRTG for monitoring.

## Getting Started:

1. Clone the repository:

   ```bash
   git clone https://github.com/minasedr/monitoring-system.git
   ```

2. Install dependencies:

   ```bash
   go get github.com/spf13/viper
   ```

3. Set up the configuration:

   - Open the `config.json` file in the project root.
   - Update the `APIURL` and `APIKey` fields with your external API's URL and API key.
   - Update the `PRTG` section with your PRTG server details, including `ServerURL`, `Username`, `Passhash`, `SensorName`, and `SensorDescription`.
   - Set the `QueryInterval` to the desired interval (in seconds) for querying the external API.

4. Run the application:
   ```bash
   go run cmd/monitoring-system/main.go
   ```

## Configuration (config.json):

```json
{
  "api_url": "https://your_api_url_here",
  "api_key": "your_api_key_here",
  "prtg": {
    "server_url": "https://your_prtg_server_url_here",
    "username": "your_prtg_username_here",
    "passhash": "your_prtg_passhash_here",
    "sensor_name": "your_sensor_name_here",
    "sensor_description": "your_sensor_description_here"
  },
  "query_interval": 3
}
```

5. Monitor the application:
   - Once the application is running, it will periodically query the external API and send data to your PRTG monitoring system with a response body and a response time.
