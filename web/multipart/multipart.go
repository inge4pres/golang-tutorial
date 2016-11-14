package mpart

import (
	"io"
	"mime/multipart"
	"net/textproto"
)

func WriteUnsignedMdnReport(out io.Writer, boundary string, text, machine []byte) error {
	outer := multipart.NewWriter(out)
	outer.SetBoundary(boundary)
	first, err := outer.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"text/plain"},
	})
	if err != nil {
		return err
	}
	if _, err := first.Write(text); err != nil {
		return err
	}

	second, err := outer.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"message/disposition-notification"},
	})
	if err != nil {
		return err
	}
	if _, err := second.Write(machine); err != nil {
		return err
	}
	return outer.Close()
}

func WriteSignedMdnReport(out io.Writer, boundary string, text, machine, signature []byte) error {
	outr := multipart.NewWriter(out)
	outr.SetBoundary(boundary)

	mdnbound := "--ReportBoundary--"
	report, err := outr.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"multipart/report; report-type=disposition-notification; boundary=\"" + mdnbound + "\""},
	})
	if err != nil {
		return err
	}
	if err = WriteUnsignedMdnReport(report, mdnbound, text, machine); err != nil {
		return err
	}
	sign, err := outr.CreatePart(textproto.MIMEHeader{
		"Content-Type":              []string{"application/pkcs7-signature; name=smime.p7s"},
		"Content-Transfer-Encoding": []string{"base64"},
		"Content-Disposition":       []string{"attachment; filename=smime.p7s"},
	})
	if _, err = sign.Write(signature); err != nil {
		return err
	}
	return outr.Close()
}
