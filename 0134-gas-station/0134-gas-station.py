class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        if sum(gas) < sum(cost):
            return -1
        
        last_diff = -1
        for i in range(len(gas)):
            if gas[i] >= cost[i] and last_diff < 0:
                if self.canComplete(i, gas, cost):
                    return i
            last_diff = gas[i] - cost[i]
            
        return -1
    
    
    def canComplete(self, start: int, gas: List[int], cost: List[int]) -> bool:
        n = len(gas)
        tank = 0
        for i in range(n):
            tank += gas[(start+i)%n] - cost[(start+i)%n]
            if tank < 0:
                return False
        return True    
        