import Data.List

digits :: Integral x => x -> [x]
digits 0 = [0]
digits x = digits (x `div` 10) ++ [x `mod` 10]

increasing :: Integral a => [a] -> Bool
increasing [] = True
increasing [x] = True
increasing (x:y:zs) = x <= y && increasing (y:zs)

-- Part 1
doubles :: Integral a => [a] -> Bool
doubles [] = False
doubles [x] = False
doubles (x:y:zs) = x == y || doubles(y:zs)

-- Part 2
doubles' :: Integral a => [a] -> Bool
doubles' xs = 1 <= length [x | x <- tuples, snd x == 2]
  where tuples = map (\xs@(x:_) -> (x, length xs)) (group xs)

lists = map digits [248345..746315]

main = do
  putStrLn $ show $ length [xs | xs <- lists, increasing xs, doubles xs]
  putStrLn $ show $ length [xs | xs <- lists, increasing xs, doubles' xs]
