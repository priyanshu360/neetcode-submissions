class MinStack {
    constructor() {
        this.min = []
        this.full = []
    }

    /**
     * @param {number} val
     * @return {void}
     */
    push(val) {
        if (this.min.length == 0) {
            this.min.push(val)
        }else{
            if (this.min[this.min.length - 1] >= val) {
                this.min.push(val)
            }
        }

        this.full.push(val)
        return null
    }

    /**
     * @return {void}
     */
    pop() {
        let val = this.full[this.full.length - 1]
        if (val == this.min[this.min.length - 1]) {
            this.min = this.min.slice(0, this.min.length - 1)
        }
        this.full = this.full.slice(0, this.full.length - 1)
        return null
    }

    /**
     * @return {number}
     */
    top() {
        return this.full[this.full.length - 1]
    }

    /**
     * @return {number}
     */
    getMin() {
        return this.min[this.min.length - 1]
    }
}
