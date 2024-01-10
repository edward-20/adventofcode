import Data.Char
import Data.Map
import qualified Data.Map as Map
-- How many steps are required?

data Node = Node { name :: String, leff :: String, righ :: String } deriving Show

-- given a character and a string, split the string up based on this character
split :: Char -> String -> [String]
split d s = [takeWhile (\c -> c /= d) s, Prelude.drop 1 $ Prelude.dropWhile (\c -> c /= d) s]

stripWhiteSpace :: String -> String
stripWhiteSpace s = (Prelude.dropWhile isSpace . reverse . Prelude.dropWhile isSpace . reverse) s

-- precondition is that the string represents a tuple
removeTupleBrackets :: String -> String
removeTupleBrackets s = Prelude.drop 1 $ Prelude.take (length s - 1) s

-- given a string that is a line representing a node
line2Node :: String -> Node
line2Node rawString =
  let 
    rawList = Prelude.map stripWhiteSpace $ Main.split '=' rawString
    rawOutgoingNodes = rawList !! 1
    outgoingNodes = Prelude.map stripWhiteSpace $ Main.split ',' $ removeTupleBrackets $ stripWhiteSpace rawOutgoingNodes
        in Node (rawList !! 0) (outgoingNodes !! 0) (outgoingNodes !! 1)

printList :: Show a => [a] -> IO ()
printList [] = return ()
printList (x:xs) = do
        print x
        printList xs

-- given the current number of steps, the current nodes, the instructions from
-- which we're up to and the tree, determine the number of steps it took to
-- get to the terminating nodes
-- followInstructions :: Int -> [Node] -> String -> Map String Node -> IO Int
-- followInstructions acc currs "" m = return acc
-- followInstructions acc currs (x:xs) m = do
--         -- get the string of the new nodes we'll be at
--         let newCurrsString = Prelude.map (\n -> if (last $ name n) == 'L' then leff n else righ n) currs :: [String]
--         -- find the nodes corresponding to the strings
--         let newCurrs = Prelude.map (\s -> findWithDefault (error "couldn't find string") s m) newCurrsString :: [Node]
--         -- check to see if the node's name end in Z
--         let isDone = all (\n -> if (last $ name n) == 'Z' then True else False) newCurrs
-- 
--         -- print acc
--         -- printList newCurrs
--         
--         if isDone then do return (acc + 1) else
--                 followInstructions (acc + 1) newCurrs xs m

followInstructions :: Int -> Node -> String -> Map String Node -> Int
followInstructions acc curr "" m = acc
followInstructions acc curr (i:is) m = 
        let newNodeString = if i == 'L' then leff curr else righ curr
            newNode = findWithDefault (error "couldn't find node in map") newNodeString m
            isFinished = if (last $ name newNode) == 'Z'then True else False
        in
        if isFinished then (acc + 1) else followInstructions (acc + 1) newNode is m

                        
createMap :: [Node] -> Map String Node -> Map String Node -- Use type (alias)
createMap [] m = m
createMap (n:ns) m = let newmap = insert (name n) n m in
        createMap ns newmap

-- find the node to go to then recurse on this node
main :: IO ()
main = do
        inp <- lines <$> readFile "input.txt"

-- get the first line; instructions
        let inst = head inp
        
-- get the list of nodes
        let nodes = Prelude.map (line2Node) $ Prelude.drop 2 inp
        let nodesMap = createMap nodes empty
         
-- create a map from the strings to the nodes
        let m1 = empty :: Map String Node
        let m2 = createMap nodes m1

-- using the tree get the nodetree (or list) that pertains to the list of
-- starts
        let starts = Prelude.filter (\n -> (last $ name n) == 'A') nodes
        printList starts
-- for each instruction we will 
        let stepsTaken = Prelude.map (\n -> followInstructions 0 n (cycle inst) m2) starts
        printList stepsTaken
        
        -- lcm :: Integral a => a -> a -> a 
        let ans = Prelude.foldl lcm 1 stepsTaken
        print ans
        -- ans <- followInstructions 0 starts (cycle inst) nodesMap
        -- print ans
