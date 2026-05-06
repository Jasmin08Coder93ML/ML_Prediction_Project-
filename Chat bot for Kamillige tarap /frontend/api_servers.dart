import 'dart:convert';
import 'package:http/http.dart' as http;

class ApiService {
  // Когда запустите сервер, здесь будет его адрес
  static const String baseUrl = 'http://localhost:5000'; 

  static Future<String> sendMessage(String message) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/chat'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'message': message}),
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        return data['reply'];
      } else {
        return 'Ошибка сервера: ${response.statusCode}';
      }
    } catch (e) {
      return 'Не удалось связаться с сервером: $e';
    }
  }
}
