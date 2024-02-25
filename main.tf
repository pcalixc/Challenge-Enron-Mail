provider "aws" {
    region = "us-east-1"  # Cambia esto a tu región preferida

}

resource "aws_instance" "example" {
  ami           = "ami-05fb0b8c1424f266b"  # Cambia esto a la AMI que desees
  instance_type = "t2.micro"

  user_data = <<-EOF
    #!/bin/bash
    yum update -y
    amazon-linux-extras install docker
    service docker start
    usermod -aG docker ec2-user

    # Instalar Git
    yum install git -y

    # Clonar el repositorio de GitHub
    clone https://pcalixc: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@github.com/pcalixc/Challenge-Enron-Mail.git /home/ec2-user/docker_project

    # Entrar al directorio del proyecto
    cd /home/ec2-user/docker_project

    # Ejecutar Docker Compose
    docker-compose up -d
  EOF
}

resource "aws_security_group" "docker" {
  name        = "docker_sg"
  description = "Allow traffic for Docker"
  
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "docker_egress" {
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.docker.id  # Asociar la regla al grupo de seguridad "docker"
}

resource "aws_security_group_rule" "docker_ingress" {
  type             = "ingress"
  from_port        = 8080
  to_port          = 8080
  protocol         = "tcp"
  cidr_blocks      = ["0.0.0.0/0"]
  security_group_id = aws_security_group.docker.id
}



output "public_ip" {
  value = aws_instance.example.public_ip
}

