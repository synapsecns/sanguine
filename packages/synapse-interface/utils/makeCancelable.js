export const makeCancelable = promise => {
  let rejectFn;

  const wrappedPromise = new Promise((resolve, reject) => {
    rejectFn = reject;

    Promise.resolve(promise)
      .then(resolve)
      .catch(reject);
  });

  wrappedPromise.cancel = () => {
    rejectFn({ canceled: true });
  };

  return wrappedPromise;
};