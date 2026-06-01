class Solution {
    /**
     * @param {number} target
     * @param {number[]} position
     * @param {number[]} speed
     * @return {number}
     */
    carFleet(target, position, speed) {
        let timetaken = new Array(position.length)

        for (let i = 0; i < position.length;  i++){
            timetaken[i] = (target - position[i]) / (speed[i] * 1.0)
        }

        for (let i = 0; i < position.length; i++) {
            for(let j = 0; j < position.length; j++) {
                if (i == j) {
                    continue
                }

                if (position[i] < position[j]) {
                    timetaken[i] = Math.max(timetaken[i], timetaken[j])
                }
            }
        }

        timetaken.sort()
        let ans = 1

        for( let i = 1; i < position.length; i++) {
            if (timetaken[i] != timetaken[i - 1]) {
                ans++;
            }
        }
        return ans;

    }
}
