output "server_ip" {
  description = "Publiczny adres IP serwera"
  value       = aws_eip.app.public_ip
}

output "app_url" {
  value = "http://${aws_eip.app.public_ip}"
}

output "grafana_url" {
  value = "http://${aws_eip.app.public_ip}:3000"
}

output "prometheus_url" {
  value = "http://${aws_eip.app.public_ip}:9090"
}
