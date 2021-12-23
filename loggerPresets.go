package littleL

import (
	"fmt"
	"time"
)

func FormatterClassic() string{
	return fmt.Sprintf("%s: $msg", time.Now().String())
}
