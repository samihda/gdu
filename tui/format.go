package tui

import (
	"fmt"

	"github.com/dundee/gdu/v5/pkg/analyze"
)

// file size constants
const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// file count constants
const (
	K int = 1e3
	M int = 1e6
	G int = 1e9
	T int = 1e12
	P int = 1e15
	E int = 1e18
)

func (ui *UI) formatFileRow(item analyze.Item) string {
	var part int

	if ui.ShowApparentSize {
		part = int(float64(item.GetSize()) / float64(item.GetParent().GetSize()) * 10.0)
	} else {
		part = int(float64(item.GetUsage()) / float64(item.GetParent().GetUsage()) * 10.0)
	}

	row := string(item.GetFlag())

	if ui.UseColors {
		row += "[#e67100::b]"
	} else {
		row += "[::b]"
	}

	if ui.ShowApparentSize {
		row += fmt.Sprintf("%15s", ui.formatSize(item.GetSize(), false, true))
	} else {
		row += fmt.Sprintf("%15s", ui.formatSize(item.GetUsage(), false, true))
	}

	row += getUsageGraph(part)

	if ui.showItemCount {
		if ui.UseColors {
			row += "[#e67100::b]"
		} else {
			row += "[::b]"
		}
		row += fmt.Sprintf("%11s ", ui.formatCount(item.GetItemCount()))
	}

	if item.IsDir() {
		if ui.UseColors {
			row += "[#3498db::b]/"
		} else {
			row += "[::b]/"
		}
	}
	row += item.GetName()
	return row
}

func (ui *UI) formatSize(size int64, reverseColor bool, transparentBg bool) string {
	var color string
	if reverseColor {
		if ui.UseColors {
			color = "[black:#2479d0:-]"
		} else {
			color = "[black:white:-]"
		}
	} else {
		if transparentBg {
			color = "[-::]"
		} else {
			color = "[white:black:-]"
		}
	}

	fsize := float64(size)

	switch {
	case fsize >= EB:
		return fmt.Sprintf("%.1f%s EiB", fsize/EB, color)
	case fsize >= PB:
		return fmt.Sprintf("%.1f%s PiB", fsize/PB, color)
	case fsize >= TB:
		return fmt.Sprintf("%.1f%s TiB", fsize/TB, color)
	case fsize >= GB:
		return fmt.Sprintf("%.1f%s GiB", fsize/GB, color)
	case fsize >= MB:
		return fmt.Sprintf("%.1f%s MiB", fsize/MB, color)
	case fsize >= KB:
		return fmt.Sprintf("%.1f%s KiB", fsize/KB, color)
	default:
		return fmt.Sprintf("%d%s B", size, color)
	}
}

func (ui *UI) formatCount(count int) string {
	row := ""
	color := "[-::]"

	switch {
	case count >= E:
		row += fmt.Sprintf("%.1f%sE", float64(count)/float64(E), color)
	case count >= P:
		row += fmt.Sprintf("%.1f%sP", float64(count)/float64(P), color)
	case count >= T:
		row += fmt.Sprintf("%.1f%sT", float64(count)/float64(T), color)
	case count >= G:
		row += fmt.Sprintf("%.1f%sG", float64(count)/float64(G), color)
	case count >= M:
		row += fmt.Sprintf("%.1f%sM", float64(count)/float64(M), color)
	case count >= K:
		row += fmt.Sprintf("%.1f%sk", float64(count)/float64(K), color)
	default:
		row += fmt.Sprintf("%d%s", count, color)
	}
	return row
}
