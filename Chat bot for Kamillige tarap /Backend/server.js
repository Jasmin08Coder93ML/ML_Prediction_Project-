const express = require('express');
const { GoogleGenerativeAI } = require("@google/generative-ai");
require('dotenv').config();

const app = express();
app.use(express.json());

// Инициализация Gemini с вашим ключом из файла .env
const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY);

app.post('/chat', async (req, res) => {
  const model = genAI.getGenerativeModel({ model: "gemini-1.5-flash" });
  const prompt = req.body.message;

  const result = await model.generateContent(prompt);
  const response = await result.response;
  res.json({ reply: response.text() });
});

const PORT = 5000;
app.listen(PORT, () => console.log(`Сервер Kamillige запущен на порту ${PORT}`));
