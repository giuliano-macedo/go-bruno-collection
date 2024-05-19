package bruno

import (
	"archive/tar"
	"encoding/json"
	"io"
	"path"
)

func writeBruTarFile(tarWriter *tar.Writer, fileName string, data []byte) error {
	hdr := &tar.Header{
		Name: fileName,
		Mode: 0600,
		Size: int64(len(data)),
	}
	if err := tarWriter.WriteHeader(hdr); err != nil {
		return err
	}
	_, err := tarWriter.Write(data)
	return err
}

func writeItemsBruTar(tarWriter *tar.Writer, items []Item, dir string) error {
	for _, item := range items {
		fname := item.Name
		if dir != "" {
			fname = path.Join(dir, fname)
		}
		if item.IsFolder() {
			if err := writeItemsBruTar(tarWriter, item.Items, fname); err != nil {
				return err
			}
			continue
		}
		if err := writeBruTarFile(tarWriter, fname+".bru", item.MarshalBru()); err != nil {
			return err
		}
	}
	return nil
}

func writeBruJsonFile(tarWriter *tar.Writer, version, name string) error {
	collectionData, err := json.Marshal(map[string]string{
		"version": version,
		"name":    name,
		"type":    "collection",
	})
	if err != nil {
		return err
	}
	if err := writeBruTarFile(tarWriter, "bruno.json", collectionData); err != nil {
		return err
	}
	return nil
}

func writeBruEnvironments(envs []Environment, tw *tar.Writer) error {
	for _, env := range envs {
		if err := writeBruTarFile(tw, path.Join("environments", env.Name+".bru"), env.MarshalBru()); err != nil {
			return err
		}
	}
	return nil
}

func CreateBruTar(collection Collection, writer io.Writer) error {
	tw := tar.NewWriter(writer)
	if err := writeBruJsonFile(tw, collection.Version, collection.Name); err != nil {
		return err
	}

	if err := writeBruEnvironments(collection.Environments, tw); err != nil {
		return err
	}

	if err := writeItemsBruTar(tw, collection.Items, ""); err != nil {
		return err
	}

	if collection.Docs != "" {
		if err := writeBruTarFile(tw, "collection.bru", collection.MarshalBru()); err != nil {
			return err
		}
	}

	return tw.Close()
}
