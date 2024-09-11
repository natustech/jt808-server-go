import socket

def tcp_client(server_ip, server_port, message):
    # Bir TCP soketi oluştur
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        try:
            # Sunucuya bağlan
            client_socket.connect((server_ip, server_port))
            print(f"Sunucuya bağlanıldı: {server_ip}:{server_port}")
            
            # Mesajı gönder
            client_socket.sendall(message)
            print(f"Mesaj gönderildi: {message}")
            
            # Yanıt al (isteğe bağlı)
            response = client_socket.recv(1024)
            print(f"Sunucudan yanıt alındı: {response}")
            while True:
                pass
            
        except Exception as e:
            print(f"Hata: {e}")

# Kullanım örneği
if __name__ == "__main__":
    server_ip = '213.254.138.190'  # Sunucunun IP adresi
    server_port = 3636       # Sunucunun dinlediği port
    message = b"""~\x07\x04\x00z\x04TP\x04R)\x00\x0f\x00\x01\x01\x00u\x00\x00\x00\x00\x00L\x00\x00\x02q\x81\x88\x01\xbc=a\x00\x00\x00\x00\x00\x00$\t\x11\x14B9\x01\x04\x00\x00\x00v0\x01\x1a1\x01\x00\xe4\x02\x012\xe5\x01\x01\xe6\x01\x00\xe7\x08\x00\x00\x00\x04\x00\x00\x00\x00\xec#\xc8Z\x9f\xde\xf4H\xd9\x16FXCN]\xcd\x88\xc3\x97\xec\x89\x82\xc6\x88@;\xfa/\x8a\xb9<w\xe6uM3\xb2\xe1\x0c\x01\x1e\x00\x01\x00\x17\r\x06-\xa5\x11\x00\xf5\x01\x01\xfb\x01\x00\xc5~"""

    tcp_client(server_ip, server_port, message)
