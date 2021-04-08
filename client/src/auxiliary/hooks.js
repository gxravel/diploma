import { useEffect, useState } from 'react';

export const useStateWithSessionStorage = sessionStorageKey => {
  const [value, setValue] = useState(
    sessionStorage.getItem(sessionStorageKey) || '',
  );
  useEffect(() => {
    sessionStorage.setItem(sessionStorageKey, value);
  }, [value, sessionStorageKey]);
  return [value, setValue];
};

export const useDebounce = (value, delay) => {
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);
    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);

  return debouncedValue;
};
