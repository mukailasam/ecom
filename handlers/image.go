package handlers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/ftsog/ecom/customerrors"
)

const (
	uploadSize = 1024 * 1024 * 1024
)

type imageData struct {
	name string
	url  string
	pID  int64
}

func (app *Handler) ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		checker, _ := app.RDsession.CheckSession(w, r)
		if !(*checker) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusUnauthorized),
					Status:  http.StatusUnauthorized,
					Message: "Unauthorize",
					Detail:  "you are not allowed to this resource, account requred",
				},
				statusCode: http.StatusUnauthorized,
			})

			return
		}

		user, err := app.RDsession.GetUserFromSession(w, r)
		if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusInternalServerError),
					Status:  http.StatusInternalServerError,
					Message: customerrors.InternalErrorMessage,
					Detail:  customerrors.InternalErrorDetail,
				},
				statusCode: http.StatusInternalServerError,
			})

			return
		}

		postIDCookie, err := r.Cookie("postID")
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
				statusCode: http.StatusBadRequest,
			})

			return
		}

		pID, err := strconv.ParseInt(postIDCookie.Value, 10, 64)
		if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusInternalServerError),
					Status:  http.StatusInternalServerError,
					Message: customerrors.InternalErrorMessage,
					Detail:  customerrors.InternalErrorDetail,
				},
				statusCode: http.StatusInternalServerError,
			})

			return
		}

		product, err := app.Db.ReadProduct(pID)
		if err != nil {
			app.Logger.ErrorLog.Println(err.Error())
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseError{
					Error:   http.StatusText(http.StatusInternalServerError),
					Status:  http.StatusInternalServerError,
					Message: customerrors.InternalErrorMessage,
					Detail:  customerrors.InternalErrorDetail,
				},
				statusCode: http.StatusInternalServerError,
			})

			return
		}

		if product.OwnerName == *user {
			err := r.ParseMultipartForm(uploadSize)
			if err != nil {
				app.Logger.ErrorLog.Println(err.Error())
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responseError{
						Error:   http.StatusText(http.StatusInternalServerError),
						Status:  http.StatusInternalServerError,
						Message: customerrors.InternalErrorMessage,
						Detail:  customerrors.InternalErrorDetail,
					},
					statusCode: http.StatusInternalServerError,
				})

				return
			}
			formData := r.MultipartForm

			files, ok := formData.File["FILE"]
			if !ok {
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responseError{
						Error:   http.StatusText(http.StatusBadRequest),
						Status:  http.StatusBadRequest,
						Message: "Invalid file upload",
					},
					statusCode: http.StatusBadRequest,
				})

				return
			}

			// accepted file format
			for k, header := range files {
				file, err := files[k].Open()
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return
				}

				buf := make([]byte, header.Size)
				_, err = file.Read(buf)
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return
				}

				ctype := http.DetectContentType(buf)
				if ctype != "image/jpeg" && ctype != "image/png" {
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusBadRequest),
							Status:  http.StatusBadRequest,
							Message: "only image allowed, no other file format supported",
						},
						statusCode: http.StatusBadRequest,
					})

					return
				}
			}

			path := "./static/images/"
			imageDir := "/images/"
			imgs := []imageData{}

			for i, _ := range files {
				img := imageData{}
				name := files[i].Filename
				url := "127.0.0.1" + ":7000" + imageDir + name
				filePath := path + name

				//file size
				for _, v := range files {
					if v.Size > uploadSize {
						JsonResponseWriter(JsonResponse{
							ResponseWriter: w,
							Data: responseError{
								Error:   http.StatusText(http.StatusBadRequest),
								Status:  http.StatusBadRequest,
								Message: "image too large",
							},
							statusCode: http.StatusBadRequest,
						})

						return
					}
				}

				f, err := os.Create(filePath)
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return
				}

				file, err := files[i].Open()
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return
				}

				_, err = io.Copy(f, file)
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return

				}

				img.name = name
				img.url = url
				img.pID = pID

				imgs = append(imgs, img)
			}

			for _, v := range imgs {
				err = app.Db.AddImage(v.name, v.url, v.pID)
				if err != nil {
					app.Logger.ErrorLog.Println(err.Error())
					JsonResponseWriter(JsonResponse{
						ResponseWriter: w,
						Data: responseError{
							Error:   http.StatusText(http.StatusInternalServerError),
							Status:  http.StatusInternalServerError,
							Message: customerrors.InternalErrorMessage,
							Detail:  customerrors.InternalErrorDetail,
						},
						statusCode: http.StatusInternalServerError,
					})

					return
				}
			}

			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responseMessage{
					Status:  http.StatusOK,
					Message: "Successfully upload image",
					Path:    "/api/product/image/create",
				},

				statusCode: http.StatusOK,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responseError{
				Error:   http.StatusForbidden,
				Status:  http.StatusText(http.StatusForbidden),
				Message: "Access denied",
			},
			statusCode: http.StatusForbidden,
		})

		return
	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responseError{
			Error:   http.StatusText(http.StatusMethodNotAllowed),
			Status:  http.StatusMethodNotAllowed,
			Message: "Wrong HTTP VERB",
			Detail:  "Make use of the appropriate HTTP verb/method",
		},
		statusCode: http.StatusMethodNotAllowed,
	})
}
