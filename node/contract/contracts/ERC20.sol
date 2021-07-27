pragma solidity >=0.4.26 <0.9.0;

interface ERC20MetaData {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
}

interface ERC20Interface {
    function totalSupply() external view returns(uint256);
    function balanceOf(address tokenOwner) external view returns(uint256);
    function allowance(address tokenOwner, address spender) external view returns(uint256);
    function transfer(address to, uint256 amount) external returns(bool);
    function approve(address spender, uint256 amount) external returns(bool);
    // function transferFrom(address froma, address to, uint256 amount) external returns(bool);

    event Transfer(address indexed froma, address indexed to, uint256 amount);
    event Approve(address indexed tokenOwner, address indexed spender, uint256 amount);
}

contract ERC20Token is ERC20Interface, ERC20MetaData {
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowance;

    uint256 private _totalSupply;
    string private _name;
    string private _symbol;

    constructor(uint256 totalSupply_, string memory name_, string memory symbol_) {
        _totalSupply = totalSupply_;
        _name = name_;
        _symbol = symbol_;
    }

    // meta data
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    function decimals() public view virtual override returns (uint8) {
        return 18;
    }

    // erc20 body
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address tokenOwner) public view virtual override returns (uint256) {
        return _balances[tokenOwner];
    }

    function allowance(address tokenOwner, address spender) public view virtual override returns (uint256) {
        return _allowance[tokenOwner][spender];
    }

    function transfer(address to, uint256 amount) public virtual override returns (bool) {
        _transfer(msg.sender, to, amount);

        return true;
    }

    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(msg.sender, spender, amount);

        return true;
    }

    function _transfer(address froma, address to, uint256 amount) internal virtual {
        require(froma != address(0), "ERC20: transfer from zero address");
        require(to != address(0), "ERC20: transfer to zero address");

        uint256 fromBalance = _balances[froma];
        require(amount <= fromBalance, "ERC20: transfer need plentiful enough balance");
        // require(amount <= _approve[froma][to], "ERC20: transfer amount should not greater than approve")


        _balances[froma] = fromBalance - amount;

        _balances[to] += amount;

        emit Transfer(froma, to, amount);
    }

    function _approve(address tokenOwner, address spender, uint256 amount) internal virtual {
        require(tokenOwner != address(0), "ERC20: transfer from zero address");
        require(spender != address(0), "ERC20: transfer to zero address");

        _allowance[tokenOwner][spender] = amount;
    }
}
