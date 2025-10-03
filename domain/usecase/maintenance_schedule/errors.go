package maintenance_schedule

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrIdIsRequired    = oops.New("Vui lòng nhập id")
	ErrCreateFailed    = oops.New("Tạo không thành công")
	ErrDeleteFailed    = oops.New("Xóa không thành công")
	ErrNotFound        = oops.New("Không tìm thấy")
	ErrUpdateFailed    = oops.New("Cập nhật không thành công")
	ErrListFailed      = oops.New("Lấy danh sách không thành công")
	ErrUnmarshalFailed = oops.New("Xử lý dữ liệu không thành công")
)
