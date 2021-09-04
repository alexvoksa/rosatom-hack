from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

import aiosmtplib

from fastapi import APIRouter

from email_sender import (
    SendMail,
    MAIL_SERVER_HOST,
    MAIL_SERVER_SSL,
    MAIL_SERVER_PORT,
    MAIL_SERVER_PASSWORD,
    MAIL_SERVER_SENDER,
)

router = APIRouter()


@router.post("/send")
async def send_email(data: SendMail):
    """Send an outgoing HTML-body email with the given parameters.

    data:
        to: str
        subject: str
        text: str
    """

    msg = MIMEMultipart()
    msg.preamble = data.subject
    msg["Subject"] = data.subject
    msg["From"] = MAIL_SERVER_SENDER
    msg["To"] = ", ".join(data.to)
    msg.attach(MIMEText(data.text, "HTML", "utf-8"))

    smtp = aiosmtplib.SMTP(
        hostname=MAIL_SERVER_HOST, port=MAIL_SERVER_PORT, use_tls=MAIL_SERVER_SSL
    )
    await smtp.starttls()
    await smtp.connect()

    await smtp.login(MAIL_SERVER_SENDER, MAIL_SERVER_PASSWORD)
    await smtp.send_message(msg)
    await smtp.quit()
