import { Card } from 'antd';

function CryptocurrencyCard(props) {
  const { currency } = props;

  const price = Math.round(currency.quote.USD.price);

  return (
    <>
      <Card
        title={
          <div className="flex items-center gap-3">
            <img
              src={`https://s2.coinmarketcap.com/static/img/coins/64x64/${currency.id}.png`}
              alt="logo"
              width="50"
            />
            <p className="text-3xl">{currency.name}</p>
          </div>
        }
        style={{ width: 300 }}
      >
        <p>Текущая цена: {price}$</p>
        <p>Card content</p>
        <p>Card content</p>
      </Card>
    </>
  );
}

export default CryptocurrencyCard;
