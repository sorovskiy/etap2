package repository

import (
	"etap2/domain"
)

func (r *Repo) GetState(application string) (domain.APIResp, error) {
	var i int

	row := r.db.QueryRow("select * from States where application = $1", application)

	resp := domain.APIResp{}
	err := row.Scan(&i, &resp.Application, &resp.Param1, &resp.Param2, &resp.Version)
	if err != nil {
		return domain.APIResp{}, err
	}
	return resp, err
}

func (r *Repo) SetState(resp domain.APIResp) error {
	_, err := r.db.Exec("insert into States (application, param1, param2, version) values ($1, $2, $3, $4)",
		resp.Application, resp.Param1, resp.Param2, resp.Version)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateState(options domain.APIResp, resp domain.APIResp) error {

	_, err := r.db.Exec("update States set param1 = $1, param2 = $2, version = $3 where application = $4",
		options.Param1, options.Param2, resp.Version+1, options.Application)
	if err != nil {
		return err
	}
	return nil

}
