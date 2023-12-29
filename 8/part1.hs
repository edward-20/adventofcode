import Data.Char
-- import Data.List.Split
-- How many steps are required?

data Node = Node { name :: String, leff :: String, righ :: String } deriving Show

-- given a character and a string, split the string up based on this character
split :: Char -> String -> [String]
split d s = [takeWhile (\c -> c /= d) s, drop 1 $ dropWhile (\c -> c /= d) s]

stripWhiteSpace :: String -> String
stripWhiteSpace s = (dropWhile isSpace . reverse . dropWhile isSpace . reverse) s

-- precondition is that the string represents a tuple
removeTupleBrackets :: String -> String
removeTupleBrackets s = drop 1 $ take (length s - 1) s

-- given a string that is a line representing a node
line2Node :: String -> Node
line2Node rawString =
  let 
    rawList = map stripWhiteSpace $ split '=' rawString
    rawOutgoingNodes = rawList !! 1
    outgoingNodes = map stripWhiteSpace $ split ',' $ removeTupleBrackets $ stripWhiteSpace rawOutgoingNodes
        in Node (rawList !! 0) (outgoingNodes !! 0) (outgoingNodes !! 1)

printList :: Show a => [a] -> IO ()
printList [] = return ()
printList (x:xs) = do
        print x
        printList xs

-- given the current number of steps and the current node
followInstructions :: Int -> Node -> String -> [Node] -> IO Int
followInstructions acc curr "" nodes = return acc
followInstructions acc curr (x:xs) nodes = do
        let new = if x == 'L' then leff curr else righ curr
        print $ show x ++ new
        let newNode = head $ dropWhile (\n -> (name n) /= new) nodes
        if new == "ZZZ" then do return (acc + 1) else
                followInstructions (acc + 1) newNode xs nodes
                        
-- find the node to go to then recurse on this node

main :: IO ()
main = do
-- <$> infix synonym for fmap. Applies function lines to Functor (monad IO
-- String) resultant from readFile "input.txt"
        inp <- lines <$> readFile "input.txt"

-- print the input
-- printList inp

-- get the first line; instructions
        let inst = head inp
-- get each node
        
        let nodes = map (line2Node) $ drop 2 inp
        let start = head $ dropWhile (\n -> name n /= "AAA") nodes
        ans <- followInstructions 0 start inst (cycle nodes)
        print ans
