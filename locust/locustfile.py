from locust import HttpUser, task, between

class WebsiteUser(HttpUser):
    wait_time = between(4, 10)  # Wait between requests
    host = "https://ausopen.com"  # Target URL

    @task
    def get_homepage(self):
        self.client.get("/")
