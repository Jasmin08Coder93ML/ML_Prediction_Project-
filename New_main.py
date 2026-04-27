import telebot
from telebot import types

TOKEN = 'ВАШ_ТОКЕН_ТУТ'
bot = telebot.TeleBot(TOKEN)

# Текст на туркменском (ваш оригинал)
text_tm = """Gelejegiňize Inwestisiýa Ediň: Sanly Bilim we Daşary Ýurt Dilleri! 🚀
... (Gelejegiňize Inwestisiýa Ediň: Sanly Bilim we Daşary Ýurt Dilleri! 🚀
​Siz özüňize ýa-da çagalaryňyza iň döwrebap bilimleri bermek isleýärmiňiz? Biz size iň amatly şertlerde, öýden çykman online formatda professional kurslary hödürleýäris!
​📚 Online Dil Kurslary (Rus we Iňlis dilleri):
​Ähli ýaşdakylar üçin: 5 ýaşdan başlap çagalar we uly ýaşlylar.
​Giriş synaglary: Ýokary okuw jaýlaryna girmek isleýän talyplar üçin ýörite taýýarlyk.
Fizika, Matematika, Informatika.
​Şertlerimiz: Hepdede 3 gün, 2 sagatdan (Jemi 3 aýlyk kurs).
​Bahasy: Aýyna bary-ýogy 200 manat (1 mln).​Wagtyňyza tygşytlylyk — islendik ýerden online gatnaşmak mümkinçiligi.
​Amatly baha we ýokary netije.
​⚠️ Üns beriň! Diňleýjileriň sany çäkli bolup, ýerlerimiz çalt dolýar. Gelejegiň hünärmeni bolmaga häzirden howlugyň!
​📲 Has giňişleýin maglumat we ýazylmak üçin:
@aysulgun77

💻 IT Kurslary – Iýun aýyndan başlangyç alýar!
Tehnologiýa dünýäsiniň iň isleg bildirilýän 11 sany ugury boýunça hünärmen boluň:
​QA Automation & Manual (Programma üpjünçiligini barlamak)
​Mobile & iOS Development (Programma döretmek)
​Full Stack Developer (Frontend & Backend)
​Machine Learning & DevOps
​Python, Java, JavaScript Developer
​Cybersecurity (Kiberhowpsuzlyk)
​✨ Nämüçin bizi saýlamaly?
​Dünýä standartlaryna laýyk okuw meýilnamasy.
​Wagtyňyza tygşytlylyk — islendik ýerden online gatnaşmak mümkinçiligi.) ..."""

# Здесь добавьте переменные text_tm и text_ru и text_en с переводами

@bot.message_handler(commands=['start'])
def start(message):
    # Отправляем рекламный текст
    bot.send_message(message.chat.id, text_tm)
    # bot.send_message(message.chat.id, text_ru) # Раскомментируйте, когда добавите перевод
    
    # Создаем кнопки выбора направления
    markup = types.InlineKeyboardMarkup(row_width=1)
    btn1 = types.InlineKeyboardButton("Online Dil Kurslary 📚", callback_data="lang_courses")
    btn2 = types.InlineKeyboardButton("IT Kurslary (Iýun) 💻", callback_data="it_courses")
    btn3 = types.InlineKeyboardButton("ML Prediksiýa (Test) 🤖", callback_data="ml_test")
    
    markup.add(btn1, btn2, btn3)
    bot.send_message(message.chat.id, "Haýsy ugur size gyzykly? / Какое направление вам интересно?", reply_markup=markup)

@bot.callback_query_handler(func=lambda call: True)
def callback_inline(call):
    if call.data == "lang_courses":
        # Показываем модули для языков
        text = "📚 **Dil kurslarynyň programmasy:**\n1. Gramatika\n2. Gepleşik\n3. Ýazuw..."
        bot.send_message(call.message.chat.id, text)
        
    elif call.data == "it_courses":
        # Показываем 11 направлений IT
        markup = types.InlineKeyboardMarkup()
        markup.add(types.InlineKeyboardButton("Python Developer", callback_data="it_python"))
        markup.add(types.InlineKeyboardButton("Cybersecurity", callback_data="it_cyber"))
        # Добавьте остальные кнопки...
        bot.send_message(call.message.chat.id, "IT ugurlaryny saýlaň:", reply_markup=markup)

    elif call.data == "ml_test":
        # Запускаем ваш старый функционал с моделью .pkl
        bot.send_message(call.message.chat.id, "Введите стаж клиента:")
        # Тут нужно вызвать логику предсказания, которую мы писали раньше
Gel
