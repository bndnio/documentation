require 'rubygems'
require 'dogapi'

api_key = '<YOUR_API_KEY>'
app_key = '<YOUR_APP_KEY>'

dog = Dogapi::Client.new(api_key, app_key)

dog.list_aws_logs_integrations
