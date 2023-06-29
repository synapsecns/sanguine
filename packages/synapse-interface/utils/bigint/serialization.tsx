import { serialize, deserialize } from 'wagmi';

export const serializeBalance = (balanceArray) => {
  return balanceArray.map(item => ({
    ...item,
    balance: item.balance !== undefined ? serialize({ value: item.balance }) : item.balance,
  }));
};

export const deserializeBalance = (balanceArray) => {
    return balanceArray.map(item => {
    //   console.log('Balance before parsing: ', item.balance);

      let balance;
      try {
        balance = deserialize((item.balance)).value;
      } catch (err) {
        console.error('Failed to parse and deserialize balance: ', err);
        balance = item.balance;  // or set it to a default value
      }

      return {
        ...item,
        balance,
      };
    });
  };
