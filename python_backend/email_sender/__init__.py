"""Initial module with data to proceed email request"""
import os

from pydantic import BaseModel

MAIL_SERVER_HOST = os.environ.get("MAIL_SERVER_HOST")
MAIL_SERVER_SENDER = os.environ.get("MAIL_SERVER_SENDER")
MAIL_SERVER_PASSWORD = os.environ.get("MAIL_SERVER_PASSWORD")
MAIL_SERVER_PORT = os.environ.get("MAIL_SERVER_PORT")
MAIL_SERVER_SSL = False


class SendMail(BaseModel):
    """Class to handle post
    request to send email"""

    to: str
    subject: str
    text: str
