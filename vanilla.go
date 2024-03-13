package main

func (v *uploadData) vanilla_upload() (err error) {
	fi, err := v.fp.Stat()

	if err != nil {
		return
	}

	err = tusOptions(v.url)

	if err != nil {
		sbox.AddLine("Server failure")
		return err
	}

	sbox.AddLine("Server exam. passed")

	location, err := tusPost(v.url, fi.Size(), v.metadata)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	sbox.AddLine("Location suuceefully aquired. \nUpload started")

	err = tusPatch(v.url, location, v.fp, fi.Size(), v.uploadFilename)

	if err != nil {
		sbox.AddLine("Upload failure")
		return err
	}

	sbox.AddLine("Upload succeeded")

	return nil
}
