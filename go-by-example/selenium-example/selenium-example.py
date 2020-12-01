import unittest
from selenium import webdriver
import time


class InstagramBot(unittest.TestCase):
    def setUp(self):
        self.driver = webdriver.Chrome("C:/chromedriver.exe")
    
    def test_instagram_login(self):
        driver = self.driver
        driver.get("https://instagram.com/")
        time.sleep(10)
        

    def tearDown(self):
        self.driver.close()



if __name__ == "__main__":
    unittest.main()