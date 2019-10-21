from datadog import initialize, api

options = {
    'api_key': '<DD_API_KEY>',
    'app_key': '<DD_APP_KEY>'
}

initialize(**options)

api.Gcp.update(
    project_id="<GCP_PROJECT_ID>",
    client_email="<GCP_CLIENT_EMAIL>",
    automute=True,
    host_filters="<NEW>:<FILTERS>"
)
