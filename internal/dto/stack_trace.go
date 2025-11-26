package dto

import "strconv"

type StackTrace struct {
	Functions []Function `json:"functions"`
}

type Function struct {
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func (st *StackTrace) Write(w *Writer) error {
	w.Descend()
	defer w.Ascend()
	for _, function := range st.Functions {
		w.AddNewline()
		_, err := w.Write([]byte(function.Name))
		if err != nil {
			return err
		}

		err = func() error {
			w.Descend()
			defer w.Ascend()

			w.AddNewline()

			_, err = w.Write([]byte(function.File + ":" + strconv.Itoa(function.Line)))
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
