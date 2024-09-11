package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/import_voters"
)

type ImportVotersPresenter struct {
	ImportPtr *common_ptr.ImportPresenter
}

func BuildImportVotersPresenter() *ImportVotersPresenter {
	return &ImportVotersPresenter{ImportPtr: common_ptr.BuildImportPresenter()}
}

func (p *ImportVotersPresenter) Show(result *import_voters.ImportVotersOutput, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError()
	}

	fc := p.ImportPtr.Format(result.Import)

	return common_adapters.OutputPort{StatusCode: http.StatusAccepted, Data: fc}
}

func (p *ImportVotersPresenter) formatError() common_adapters.OutputPort {
	return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
}
