package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignID(input GetTransactionsInput) ([]Transaction, error) {
	transactions, err := s.repository.GetByCampaignID(input.ID, input.User.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
