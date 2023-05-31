#Алфавит на Английском
go build 

# Шифруем
# Сообшение тут: "The quick brown fox jumps over the lazy dog" 
# Ключ: 20
echo The quick brown fox jumps over the lazy dog | ./frequency_analysis -e 20 


# Дешифрование
echo Dro aesmu lbygx pyh tewzc yfob dro vkji nyq | ./frequency_analysis -e -20       


# Частоты букв
echo The quick brown fox jumps over the lazy dog | ./frequency_analysis -f           


# Ключ тут 20
echo The quick brown fox jumps over the lazy dog | ./frequency_analysis -e 20 > msg 

cat msg


# Взламываем и находим сам ключ
./frequency_analysis -c < msg
