from unidecode import unidecode
import re

def isPalindrome(phrase):
    phrase = cleanPhrase(phrase)
    backwards = len(phrase) - 1
    for i in range(0, len(phrase) - 1):	
        if unidecode(phrase[i]) == unidecode(phrase[backwards]):
            i += 1
            backwards -= 1
        else:
            return False
    return True
     
def cleanPhrase(original):
    clean = re.sub(r'[\s,.;:\"\']', "", original)
    clean = clean.lower()
    return clean
    

print("is \"ababa\" a palindrome?")
print(isPalindrome("ababa"))

print("is \"La ruta nos aportó otro paso natural\" a palindrome?")
print(isPalindrome("La ruta nos aportó otro paso natural."))
