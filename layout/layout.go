package layout

import (
	"github.com/marks12/gothamel/tag"
)

func Doc() tag.Html {

	v := tag.Html{
		Head: tag.Head{Title: "Some title"},
	}

	return v

}
