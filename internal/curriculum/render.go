package curriculum

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/paulcalimache/go-curriculum/internal/models"
)

type CurriculumFile struct {
	file bytes.Buffer
}

func RenderTemplate(tmplName string, cv *models.CV) (*CurriculumFile, error) {
	tmplFile := tmplName + ".html"
	tmplPath := "templates/" + tmplName + "/" + tmplFile
	stylePath := "templates/" + tmplName + "/" + "style.html"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplPath, stylePath)
	if err != nil {
		return nil, err
	}

	c := &CurriculumFile{}

	err = tmpl.ExecuteTemplate(&c.file, "classic.html", &cv)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c CurriculumFile) SaveAsHTML(output string) error {
	err := os.MkdirAll(output, os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(output, "curriculum.html"), c.file.Bytes(), 0644)
}

func (c CurriculumFile) SaveAsPDF(output string) error {
	err := os.MkdirAll(output, os.ModePerm)
	if err != nil {
		return err
	}

	ctx, cancelCtx := chromedp.NewContext(context.Background(), chromedp.WithLogf(slog.Info))
	defer cancelCtx()

	var wg sync.WaitGroup
	wg.Add(1)

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					wg.Done()
					// remove event listener
					cancel()
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, c.file.String()).Do(ctx)
		}),
		// wait for the page.EventLoadEventFired
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.
				PrintToPDF().
				WithMarginBottom(0).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginLeft(0).
				WithPaperHeight(11.67).
				WithPaperWidth(8.27).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				slog.Debug("test 5")
				return err
			}
			return os.WriteFile(filepath.Join(output, "curriculum.pdf"), buf, 0644)
		}),
	); err != nil {
		return err
	}
	return nil
}
